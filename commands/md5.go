package commands

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/spf13/cobra"
)

func Md5CommandHandler(ctx *CommandContext) {
	ctx.Result = getMD5Hash(ctx.Input)
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func NewMd5Command(ctx *CommandContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "md5",
		Short: "Creates an MD5 hash",
		Run: func(cmd *cobra.Command, args []string) {
			Md5CommandHandler(ctx)
		},
	}

	flags := cmd.Flags()
	flags.SetInterspersed(false)

	return cmd
}
