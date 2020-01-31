package appcontext

//List of consts containing the names of the available componentes in the Application Context - appcontext.Current
const (
	Config           = "Config"
	ConfigRepository = "ConfigRepository"
	RootCmd          = "RootCmd"
)

//Component is the Base interface for all Components
type Component interface{}

//ComponentInitializerFunction specifies a function for lazily initializing a component
type ComponentInitializerFunction func() Component

//ComponentInfo holds the function to lazy initialize the component and the instance created following the singleton pattern
type ComponentInfo struct {
	Initializer ComponentInitializerFunction
	Instance    Component
}

//Get s the instance. If it is not created, creates and stores it to the next calls
func (componentInfo *ComponentInfo) Get() Component {
	if componentInfo.Instance == nil {
		componentInfo.Instance = componentInfo.Initializer()
	}

	return componentInfo.Instance
}

//ApplicationContext is the type defining a map of Components
type ApplicationContext struct {
	components map[string]*ComponentInfo
}

//Current keeps all components available, initialized in the application startup
var Current ApplicationContext

//Add a component By Name
func (applicationContext *ApplicationContext) Add(componentName string, componentInitializerFunction ComponentInitializerFunction) {
	applicationContext.components[componentName] = &ComponentInfo{Initializer: componentInitializerFunction}
}

//Get a component By Name
func (applicationContext *ApplicationContext) Get(componentName string) Component {
	if applicationContext.components[componentName] == nil {
		return nil
	}
	return applicationContext.components[componentName].Get()
}

//Delete a component By Name
func (applicationContext *ApplicationContext) Delete(componentName string) {
	delete(applicationContext.components, componentName)
}

//Count returns the count of components registered
func (applicationContext *ApplicationContext) Count() int {
	return len(applicationContext.components)
}

//CreateApplicationContext creates a new ApplicationContext instance
func CreateApplicationContext() ApplicationContext {
	return ApplicationContext{components: make(map[string]*ComponentInfo)}
}

func init() {
	Current = CreateApplicationContext()
}
