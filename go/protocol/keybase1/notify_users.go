// Auto-generated types and interfaces using avdl-compiler v1.4.1 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/notify_users.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type UserChangedArg struct {
	Uid UID `codec:"uid" json:"uid"`
}

type PasswordChangedArg struct {
}

type IdentifyUpdateArg struct {
	OkUsernames     []string `codec:"okUsernames" json:"okUsernames"`
	BrokenUsernames []string `codec:"brokenUsernames" json:"brokenUsernames"`
}

type NotifyUsersInterface interface {
	UserChanged(context.Context, UID) error
	PasswordChanged(context.Context) error
	IdentifyUpdate(context.Context, IdentifyUpdateArg) error
}

func NotifyUsersProtocol(i NotifyUsersInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.NotifyUsers",
		Methods: map[string]rpc.ServeHandlerDescription{
			"userChanged": {
				MakeArg: func() interface{} {
					var ret [1]UserChangedArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]UserChangedArg)
					if !ok {
						err = rpc.NewTypeError((*[1]UserChangedArg)(nil), args)
						return
					}
					err = i.UserChanged(ctx, typedArgs[0].Uid)
					return
				},
			},
			"passwordChanged": {
				MakeArg: func() interface{} {
					var ret [1]PasswordChangedArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					err = i.PasswordChanged(ctx)
					return
				},
			},
			"identifyUpdate": {
				MakeArg: func() interface{} {
					var ret [1]IdentifyUpdateArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]IdentifyUpdateArg)
					if !ok {
						err = rpc.NewTypeError((*[1]IdentifyUpdateArg)(nil), args)
						return
					}
					err = i.IdentifyUpdate(ctx, typedArgs[0])
					return
				},
			},
		},
	}
}

type NotifyUsersClient struct {
	Cli rpc.GenericClient
}

func (c NotifyUsersClient) UserChanged(ctx context.Context, uid UID) (err error) {
	__arg := UserChangedArg{Uid: uid}
	err = c.Cli.Notify(ctx, "keybase.1.NotifyUsers.userChanged", []interface{}{__arg})
	return
}

func (c NotifyUsersClient) PasswordChanged(ctx context.Context) (err error) {
	err = c.Cli.Notify(ctx, "keybase.1.NotifyUsers.passwordChanged", []interface{}{PasswordChangedArg{}})
	return
}

func (c NotifyUsersClient) IdentifyUpdate(ctx context.Context, __arg IdentifyUpdateArg) (err error) {
	err = c.Cli.Notify(ctx, "keybase.1.NotifyUsers.identifyUpdate", []interface{}{__arg})
	return
}
