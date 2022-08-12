/*
 * Copyright (c) 2022. Hiroki Okui
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 */

package cmd

import (
	"context"
	"github.com/hrk091/cobra-test/pkg/logger"
	"github.com/hrk091/cobra-test/pkg/sample"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := newHelloCfg(cmd)
		logger.Setup(cfg.Devel, cfg.Verbose)
		ctx := logger.WithLogger(context.Background(), logger.NewLogger())
		cobra.CheckErr(sample.RunHello(ctx, cfg))
	},
}

func init() {
	helloCmd.Flags().StringP("msg", "m", "World", "Greet to")
	mustBindToViper(helloCmd)
}

func newHelloCfg(cmd *cobra.Command) sample.HelloCfg {
	msg := viper.GetString("msg")

	return sample.HelloCfg{
		RootCfg: newRootCfg(cmd),
		Msg:     msg,
	}
}
