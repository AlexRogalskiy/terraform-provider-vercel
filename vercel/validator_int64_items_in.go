package vercel

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func int64ItemsIn(items ...int64) validatorInt64ItemsIn {
	itemMap := map[int64]struct{}{}
	for _, i := range items {
		itemMap[i] = struct{}{}
	}
	return validatorInt64ItemsIn{
		Items: itemMap,
	}
}

type validatorInt64ItemsIn struct {
	Items map[int64]struct{}
}

func (v validatorInt64ItemsIn) keys() (out []string) {
	for k := range v.Items {
		out = append(out, strconv.Itoa(int(k)))
	}
	return
}

func (v validatorInt64ItemsIn) Description(ctx context.Context) string {
	return fmt.Sprintf("Item must be one of %s", strings.Join(v.keys(), " "))
}
func (v validatorInt64ItemsIn) MarkdownDescription(ctx context.Context) string {
	return fmt.Sprintf("Item must be one of `%s`", strings.Join(v.keys(), "` `"))
}

func (v validatorInt64ItemsIn) Validate(ctx context.Context, req tfsdk.ValidateAttributeRequest, resp *tfsdk.ValidateAttributeResponse) {
	var item types.Int64
	diags := tfsdk.ValueAs(ctx, req.AttributeConfig, &item)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}
	if item.Unknown || item.Null {
		return
	}

	if _, ok := v.Items[item.Value]; !ok {
		resp.Diagnostics.AddAttributeError(
			req.AttributePath,
			"Invalid value provided",
			fmt.Sprintf("Item must be one of %s, got: %d.", strings.Join(v.keys(), " "), item.Value),
		)
		return
	}
}
