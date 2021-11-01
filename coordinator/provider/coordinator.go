package provider

import (
	"context"
	"net"

	"github.com/jackc/pgproto3/v2"
	"github.com/pg-sharding/spqr/coordinator"
	"github.com/pg-sharding/spqr/pkg/client"
	"github.com/pg-sharding/spqr/pkg/config"
	"github.com/pg-sharding/spqr/pkg/conn"
	"github.com/pg-sharding/spqr/pkg/models/kr"
	"github.com/pg-sharding/spqr/pkg/models/shrule"
	"github.com/pg-sharding/spqr/qdb/qdb"
	"github.com/pg-sharding/spqr/router/grpcclient"
	router "github.com/pg-sharding/spqr/router/pkg"
	"github.com/pg-sharding/spqr/router/pkg/rrouter"
	routerproto "github.com/pg-sharding/spqr/router/protos"
	spqrparser "github.com/pg-sharding/spqr/yacc/console"
	"github.com/wal-g/tracelog"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
)

type routerconn struct {
	routerproto.KeyRangeServiceClient
	addr string
	id   string
}

func (r *routerconn) Addr() string {
	return r.addr
}

func (r *routerconn) ID() string {
	return r.id
}

var _ router.Router = &routerconn{}

func DialRouter(r router.Router) (*grpc.ClientConn, error) {
	return grpcclient.Dial(r.Addr())
}

type qdbCoordinator struct {
	client.InteractRunner
	coordinator.Coordinator
	db qdb.QrouterDB
}

func (d *qdbCoordinator) ListShardingRules() ([]*shrule.ShardingRule, error) {
	return d.db.ListShardingRules()
}

func (d *qdbCoordinator) AddShardingRule(rule *shrule.ShardingRule) error {
	resp, err := d.db.ListRouters()
	if err != nil {
		return err
	}
	tracelog.InfoLogger.Printf("routers %v", resp)
	for _, r := range resp {
		cc, err := DialRouter(r)

		tracelog.InfoLogger.Printf("dialing router %v, err %w", r, err)
		if err != nil {
			return err
		}

		cl := routerproto.NewShardingRulesServiceClient(cc)
		resp, err := cl.AddShardingRules(context.TODO(), &routerproto.AddShardingRuleRequest{
			Rules: []*routerproto.ShardingRule{{Columns: rule.Columns()}},
		})

		if err != nil {
			return err
		}

		tracelog.InfoLogger.Printf("got resp %v", resp)
	}

	return nil
}

func (d *qdbCoordinator) AddLocalTable(tname string) error {
	panic("implement me")
}

func (d *qdbCoordinator) AddKeyRange(ctx context.Context, keyRange *kr.KeyRange) error {
	resp, err := d.db.ListRouters()
	if err != nil {
		return err
	}
	tracelog.InfoLogger.Printf("routers %v", resp)
	for _, r := range resp {
		cc, err := DialRouter(r)

		tracelog.InfoLogger.Printf("dialing router %v, err %w", r, err)
		if err != nil {
			return err
		}

		cl := routerproto.NewKeyRangeServiceClient(cc)
		resp, err := cl.AddKeyRange(context.TODO(), &routerproto.AddKeyRangeRequest{
			KeyRange: keyRange.ToProto(),
		})

		if err != nil {
			return err
		}

		tracelog.InfoLogger.Printf("got resp %v", resp)
	}

	return nil
}

var _ coordinator.Coordinator = &qdbCoordinator{}

func NewCoordinator(db qdb.QrouterDB) *qdbCoordinator {
	return &qdbCoordinator{
		db: db,
	}
}

func (d *qdbCoordinator) RegisterRouter(r *qdb.Router) error {

	tracelog.InfoLogger.Printf("register router %v %v", r.Addr(), r.ID())

	return d.db.AddRouter(r)
}

func (d *qdbCoordinator) ProcClient(ctx context.Context, nconn net.Conn) error {
	cl := rrouter.NewPsqlClient(nconn)

	err := cl.Init(nil, config.SSLMODEDISABLE)

	if err != nil {
		return err
	}

	tracelog.InfoLogger.Printf("initialized client connection %s-%s\n", cl.Usr(), cl.DB())

	if err := cl.AssignRule(&config.FRRule{
		AuthRule: config.AuthRule{
			Method: config.AuthOK,
		},
	}); err != nil {
		return err
	}

	if err := cl.Auth(); err != nil {
		return err
	}
	tracelog.InfoLogger.Printf("client auth OK")

	for {
		msg, err := cl.Receive()
		if err != nil {
			tracelog.ErrorLogger.Printf("failed to received msg %w", err)
			return err
		}

		tracelog.InfoLogger.Printf("received msg %v", msg)

		switch v := msg.(type) {
		case *pgproto3.Query:

			tstmt, err := spqrparser.Parse(v.String)
			if err != nil {
				return err
			}

			tracelog.InfoLogger.Printf("parsed %T", v.String, tstmt)

			if err := func() error {
				switch stmt := tstmt.(type) {
				case *spqrparser.ShardingColumn:
					err := d.AddShardingRule(shrule.NewShardingRule([]string{stmt.ColName}))
					if err != nil {
						cl.ReplyErr(err.Error())
						tracelog.ErrorLogger.PrintError(err)
						return err
					}

					return nil
				case *spqrparser.RegisterRouter:
					err := d.RegisterRouter(qdb.NewRouter(stmt.Addr, stmt.ID))
					if err != nil {
						cl.ReplyErr(err.Error())
						tracelog.ErrorLogger.PrintError(err)
						return err
					}

					return nil
				case *spqrparser.KeyRange:
					if err := d.AddKeyRange(ctx,
						kr.KeyRangeFromSQL(stmt),
					); err != nil {
						cl.ReplyErr(err.Error())
						tracelog.ErrorLogger.PrintError(err)
						return err
					}

					return nil
				case *spqrparser.Show:

				default:
					return xerrors.New("failed to proc")
				}
				return nil
			}(); err != nil {
				_ = cl.ReplyErr(err.Error())
			} else {

				if err := func() error {
					for _, msg := range []pgproto3.BackendMessage{
						&pgproto3.RowDescription{Fields: []pgproto3.FieldDescription{
							{
								Name:                 []byte("coordinator"),
								TableOID:             0,
								TableAttributeNumber: 0,
								DataTypeOID:          25,
								DataTypeSize:         -1,
								TypeModifier:         -1,
								Format:               0,
							},
						}},
						&pgproto3.DataRow{Values: [][]byte{[]byte("query ok")}},
						&pgproto3.CommandComplete{CommandTag: []byte("SELECT 1")},
						&pgproto3.ReadyForQuery{
							TxStatus: conn.TXREL,
						},
					} {
						if err := cl.Send(msg); err != nil {
							return err
						}
					}

					return nil
				}(); err != nil {
					tracelog.ErrorLogger.PrintError(err)
				}
			}

		default:
		}
	}
}
