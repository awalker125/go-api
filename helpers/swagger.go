package helpers

import (
	"fmt"
	"net/url"
)

const SwaggerBeforeScriptJs = `const UrlMutatorPlugin = (system) => ({
	rootInjects: {
	  setScheme: (scheme) => {
		const jsonSpec = system.getState().toJSON().spec.json;
		const schemes = Array.isArray(scheme) ? scheme : [scheme];
		const newJsonSpec = Object.assign({}, jsonSpec, { schemes });
  
		return system.specActions.updateJsonSpec(newJsonSpec);
	  },
	  setHost: (host) => {
		const jsonSpec = system.getState().toJSON().spec.json;
		const newJsonSpec = Object.assign({}, jsonSpec, { host });
  
		return system.specActions.updateJsonSpec(newJsonSpec);
	  },
	  setBasePath: (basePath) => {
		const jsonSpec = system.getState().toJSON().spec.json;
		const newJsonSpec = Object.assign({}, jsonSpec, { basePath });
  
		return system.specActions.updateJsonSpec(newJsonSpec);
	  }
	}
  });`

func UrlMutatorPlugin() []string {
	return []string{"UrlMutatorPlugin"}
}

func SwaggerUIConfig(uri *url.URL) map[string]string {
	return map[string]string{
		"onComplete": fmt.Sprintf(`() => {
			window.ui.setScheme('%s');
			window.ui.setHost('%s');
			window.ui.setBasePath('%s');
		}`, uri.Scheme, uri.Host, uri.Path),
	}
}
