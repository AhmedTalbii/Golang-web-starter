
# Go Web Project Structure with Go Templates (Personal Preference)

This repository demonstrates a clean and practical project structure for building web applications in Go, using Go's standard library templates (`html/template`) for the frontend and Go for backend logic.

> **Note:** This structure reflects my personal preference and workflow. It works well for many web projects, but you might want to adapt or extend it depending on your specific requirements.

---

## Project Layout

```
.
├── main.go                      # Application entry point
├── server/
│   └── server.go                # HTTP server setup and start
├── config/
│   ├── port.go                  # Server port and configuration constants
│   └── paths.go                 # Paths to templates and static assets
├── routes/
│   └── router.go                # HTTP route definitions and handlers
├── controllers/
│   ├── render/                  # Template rendering logic
│   │   └── renderPage.go
│   ├── homeController.go        # Controller for home page
│   ├── aboutController.go       # Controller for about page
│   ├── contactController.go     # Controller for contact page
│   ├── helpController.go        # Controller for help page
│   └── errorController.go       # Controller for error pages
├── models/                      # Data models and structs
├── views/
│   ├── pages/                   # Full page templates (layout + content)
│   │   ├── layout.html          # Main layout template
│   │   ├── home.html
│   │   ├── about.html
│   │   ├── contact.html
│   │   ├── help.html
│   │   └── error.html
│   └── sections/                # Reusable template sections (partials)
│       └── navbar.html          # Navigation bar partial
├── go.mod                       # Go module file
└── go.sum                       # Go module checksums
```

---

## Why This Structure?

- **Separation of Concerns:** Controllers handle HTTP requests and business logic, views handle the UI templates, and routes define URL mappings.
- **Reusable Templates:** Using partials like `navbar.html` helps avoid duplication and makes templates easier to maintain.
- **Scalability:** Adding new pages or components is straightforward—just add a new controller and page template.
- **Config Centralization:** Server port and path constants are kept in the `config` folder, making global changes easier.
- **Readability & Maintainability:** Clear folder naming and file organization make the codebase approachable for you or collaborators.

---

## Tips to Improve and Extend

- **Static Assets (CSS, JS, images):**  
  Create a `/static/` folder at the root (e.g., `./static/css/`, `./static/js/`) and serve it using `http.FileServer` in your server setup.  
  Update your templates to link stylesheets like:  
  ```html
  <link rel="stylesheet" href="/static/css/style.css">
  ```

- **Middleware:**  
  When you need logging, authentication, recovery, or compression, add a `middlewares/` package and wrap your handlers accordingly.

- **Models:**  
  Organize your data structures, database access, and business logic in the `models/` package. This keeps your application’s core data logic clean and reusable.

- **Template Caching:**  
  For performance, parse and cache templates once during app start instead of parsing on every request.

- **Environment Variables:**  
  Use `.env` or OS environment variables for configuration (port, database URLs) rather than hardcoding in `config/`.

- **Modularization:**  
  As your project grows, consider splitting controllers and models into feature-specific packages (e.g., `user/`, `product/`).

- **Error Handling:**  
  Centralize error handling and create custom error pages.

---

## How to Use This Structure

1. Put your Go source files as described above.
2. Put HTML templates in `views/pages/` and partials in `views/sections/`.
3. Add static assets (CSS/JS/images) in a dedicated `static/` folder.
4. Serve your static folder from Go like:  
   ```go
   mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
   ```
5. Use `RenderPage` in controllers to render templates.
6. Run your app and visit `http://localhost:<port>`.

---

## Final Thoughts

This structure is a great starting point for Go web applications that use server-side rendering with Go templates. It balances simplicity with organization, making it easy to extend and maintain.

Feel free to fork, adapt, and improve it for your own projects!

---

If you find this helpful, please ⭐ the repo and share with others looking for a clean Go web project structure!

---

**Happy coding!** 🚀
