package main

import (
	"fmt"

	"github.com/flosch/pongo2"
)

func main() {
	// Compile the template first (i. e. creating the AST)
	tpl, err := pongo2.FromString(`
<html>
  <head>
    <title>Our admins and users</title>
  </head>

  // This is a short example to give you a quick overview of pongo2's syntax. 

  {% macro user_details(is_admin=false) %}
  <div class="user_item">
    <!-- Let's indicate a user's good karma -->
    <p>This user is an admin!</p>
  </div>
  {% endmacro %}

  <body>
    <!-- Make use of the macro defined above to avoid repetitive HTML code
         since we want to use the same code for admins AND members -->

    <h1>Our admins</h1>
	{{ user_details(true) }}
	{{name}}
    <h1>Our members</h1>
  </body>
</html>
`)
	if err != nil {
		panic(err)
	}
	// Now you can render the template with the given
	// pongo2.Context how often you want to.
	out, err := tpl.Execute(pongo2.Context{"name": "florian"})
	if err != nil {
		panic(err)
	}
	fmt.Println(out) // Output: Hello Florian!
}
