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
package finance

import (
	"fmt"
	"math"

	"github.com/MakeNowJust/heredoc"
	"github.com/renato0307/canivete/pkg/iostreams"
	"github.com/spf13/cobra"
)

func NewCompoundInterestsCmd(iostreams iostreams.IOStreams) *cobra.Command {

	var compoundInterestsCmd = &cobra.Command{
		Use:   "compoundinterests",
		Short: "Calculates compound interests",
		Long: heredoc.Doc(`
			Calculates compound interests.
	
			The formula for compound interests is A = P * ((1 + r/n) ^ (n * t))
	
			Where:
	
			A = the future value of the investment/loan, including interest
			P = the principal investment amount (the initial deposit or loan amount)
			r = the annual interest rate (decimal)
			n = the number of times that interest is compounded per unit t
			t = the time the money is invested or borrowed for
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			investAmount, _ := cmd.Flags().GetInt("invest-amount")
			annualInterestRate, _ := cmd.Flags().GetFloat64("annual-interest-rate")
			compoundPeriods, _ := cmd.Flags().GetInt("compound-periods")
			time, _ := cmd.Flags().GetInt("time")

			part1 := (1 + ((annualInterestRate / 100) / float64(compoundPeriods)))
			part2 := float64(compoundPeriods * time)
			result := float64(investAmount) * math.Pow(part1, part2)
			roundedResult := math.Ceil(result*100) / 100

			_, err := fmt.Fprintln(iostreams.Out, roundedResult)

			return err
		},
	}

	compoundInterestsCmd.Flags().IntP("invest-amount", "p", 0, "the principal investment amount (the initial deposit or loan amount)")
	compoundInterestsCmd.Flags().Float64P("annual-interest-rate", "r", 0.0, "the annual interest rate (decimal, percentage)")
	compoundInterestsCmd.Flags().IntP("compound-periods", "n", 0, "the number of times that interest is compounded per unit t (e.g. 12 if monthly)")
	compoundInterestsCmd.Flags().IntP("time", "t", 0, "the time the money is invested or borrowed for (e.g. 10 years)")

	compoundInterestsCmd.MarkFlagRequired("invest-amount")
	compoundInterestsCmd.MarkFlagRequired("annual-interest-rate")
	compoundInterestsCmd.MarkFlagRequired("compound-periods")
	compoundInterestsCmd.MarkFlagRequired("time")

	return compoundInterestsCmd
}
