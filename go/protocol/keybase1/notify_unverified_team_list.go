// Auto-generated types and interfaces using avdl-compiler v1.4.1 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/notify_unverified_team_list.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type TeamListUnverifiedChangedArg struct {
	TeamName string `codec:"teamName" json:"teamName"`
}

type NotifyUnverifiedTeamListInterface interface {
	TeamListUnverifiedChanged(context.Context, string) error
}

func NotifyUnverifiedTeamListProtocol(i NotifyUnverifiedTeamListInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.NotifyUnverifiedTeamList",
		Methods: map[string]rpc.ServeHandlerDescription{
			"teamListUnverifiedChanged": {
				MakeArg: func() interface{} {
					var ret [1]TeamListUnverifiedChangedArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]TeamListUnverifiedChangedArg)
					if !ok {
						err = rpc.NewTypeError((*[1]TeamListUnverifiedChangedArg)(nil), args)
						return
					}
					err = i.TeamListUnverifiedChanged(ctx, typedArgs[0].TeamName)
					return
				},
			},
		},
	}
}

type NotifyUnverifiedTeamListClient struct {
	Cli rpc.GenericClient
}

func (c NotifyUnverifiedTeamListClient) TeamListUnverifiedChanged(ctx context.Context, teamName string) (err error) {
	__arg := TeamListUnverifiedChangedArg{TeamName: teamName}
	err = c.Cli.Notify(ctx, "keybase.1.NotifyUnverifiedTeamList.teamListUnverifiedChanged", []interface{}{__arg})
	return
}
