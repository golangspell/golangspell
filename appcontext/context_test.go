package appcontext

import (
	"testing"
)

func TestContext_Add(t *testing.T) {
	type fields struct {
		components map[string]ComponentInfo
	}
	type args struct {
		componentName string
		component     ComponentInfo
	}
	components := make(map[string]ComponentInfo)
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Add Component and Delete",
			fields: fields{components: components},
			args: args{
				componentName: Config,
				component:     ComponentInfo{Initializer: func() Component { return ApplicationContext{} }},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			context := CreateApplicationContext()
			context.Add(tt.args.componentName, tt.args.component.Initializer)
			if context.Count() == 0 {
				t.Error("Component not added")
			}
			Config :=
				context.Get(Config)
			if Config == nil {
				t.Error("Component not found")
			}
			context.Delete("Config")
			Config =
				context.Get("Config")
			if Config != nil {
				t.Error("Component not deleted")
			}
		})
	}

}
