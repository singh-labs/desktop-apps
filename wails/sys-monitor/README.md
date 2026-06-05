# README

## About

This is the official Wails React-TS template.

You can configure the project by editing `wails.json`. More information about the project settings can be found
here: https://wails.io/docs/reference/project-config

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

When you run wails dev (or wails generate module), a frontend module will be generated containing the following:

1. JavaScript bindings for all bound methods
2. TypeScript declarations for all bound methods
3. TypeScript definitions for all Go structs used as inputs or outputs by the bound methods

This makes it incredibly simple to call Go code from the frontend, using the same strongly typed datastructures.

## Building

To build a redistributable, production mode package, use `wails build`.

## Important Links 

1. https://wails.io/docs/next/howdoesitwork - How Wails works
2. https://wails.io/docs/next/guides/application-development - Application development guide
