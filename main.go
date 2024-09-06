package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		vpc, err := NewVpc(ctx, "eks-vpc", "10.10.0.0/16")
		if err != nil {
			return err
		}

		ctx.Export("vpcName", vpc.ID())
		return nil
	})
}
