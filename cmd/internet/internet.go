/*
Copyright © 2021 Renato Torres <renato.torres@pm.me>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package internet

import (
	"fmt"

	"github.com/renato0307/canivete/pkg/iostreams"
	"github.com/spf13/cobra"
)

func NewInternetCmd(iostreams iostreams.IOStreams) *cobra.Command {

	var internetCmd = &cobra.Command{
		Use:   "internet",
		Short: "Mist internet stuff",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("must specify a subcommand")
		},
	}

	internetCmd.AddCommand(NewMediumToMdCmd(iostreams))

	return internetCmd
}
