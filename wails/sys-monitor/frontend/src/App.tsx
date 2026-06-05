import { useEffect, useState } from "react";
import { Greet, GetPerson, UpdateLoginCount } from "../wailsjs/go/main/App";
import { main } from "../wailsjs/go/models";

import { Sys } from "./Sys";

import logo from "./assets/images/logo-universal.png";
import "./App.css";

type Person = main.Person | null;
// OR
// type Person = Awaited<ReturnType<typeof GetPerson>> | null;

function App() {
  const [resultText, setResultText] = useState("Please enter your name below 👇");
  const [name, setName] = useState("");
  const [person, setPerson] = useState<Person | null>(null);

  const updateName = (e: any) => setName(e.target.value);
  const updateResultText = (result: string) => setResultText(result);

  // initialise person state by calling GetPerson on component mount
  useEffect(() => {
    // OR
    // let person = new main.Person();
    GetPerson().then((result) => {
      setPerson(result);
    });
  }, []);

  function greet() {
    /**
     * The generated methods return a Promise. A successful call will result in the first return value from the Go call to be passed to the resolve handler. An unsuccessful call is when a Go method that has an error type as it's second return value, passes an error instance back to the caller. This is passed back via the reject handler. In the example above, Greet only returns a string so the JavaScript call will never reject - unless invalid data is passed to it.
     * All data types are correctly translated between Go and JavaScript. Even structs. If you return a struct from a Go call, it will be returned to your frontend as a JavaScript class.
     */
    Greet(name).then(updateResultText);
  }

  // LogConsoleInfo("App component rendered 🙂!");

  function updateLoginCount() {
    if (!person) {
      return;
    }
    // UpdateLoginCount accepts a Person struct, updates the login count and returns the updated struct back to the frontend. This is an example of how you can work with complex data types in your Go calls. You can pass structs, slices, maps etc. to and from your Go calls and they will be correctly translated to their JavaScript equivalents.
    UpdateLoginCount(person).then((result) => {
      setPerson(result);
    });
  }

  return (
    <div id="App">
      <img
        src={logo}
        id="logo"
        alt="logo"
      />

      <div className="container">
        <div
          id="input"
          className="input-box"
        >
          <h3>Simple log</h3>
          <div
            id="result"
            className="result"
          >
            {resultText}
          </div>
          <input
            id="name"
            className="input"
            onChange={updateName}
            autoComplete="off"
            name="input"
            type="text"
          />
          <button
            className="btn"
            onClick={greet}
          >
            Greet
          </button>
        </div>

        <div className="json-result">
          <h3>GO {"<->"} JavaScript</h3>
          <div className="code-block">
            <code>{JSON.stringify(person, null, 2)}</code>
          </div>

          <button
            className="btn"
            onClick={updateLoginCount}
          >
            Update login count - (Backend)
          </button>
        </div>
        <div>
          <Sys />
        </div>
      </div>
    </div>
  );
}

export default App;
