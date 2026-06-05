import { useState } from "react";
import logo from "./assets/images/logo-universal.png";
import "./App.css";
import { Greet, LogConsoleInfo } from "../wailsjs/go/main/App";

function App() {
  const [resultText, setResultText] = useState("Please enter your name below 👇");
  const [name, setName] = useState("");
  const updateName = (e: any) => setName(e.target.value);
  const updateResultText = (result: string) => setResultText(result);

  function greet() {
    /**
     * The generated methods return a Promise. A successful call will result in the first return value from the Go call to be passed to the resolve handler. An unsuccessful call is when a Go method that has an error type as it's second return value, passes an error instance back to the caller. This is passed back via the reject handler. In the example above, Greet only returns a string so the JavaScript call will never reject - unless invalid data is passed to it.
     * All data types are correctly translated between Go and JavaScript. Even structs. If you return a struct from a Go call, it will be returned to your frontend as a JavaScript class.
     */
    Greet(name).then(updateResultText);
  }

  LogConsoleInfo("App component rendered 🙂!");

  return (
    <div id="App">
      <img
        src={logo}
        id="logo"
        alt="logo"
      />
      <div
        id="result"
        className="result"
      >
        {resultText}
      </div>
      <div
        id="input"
        className="input-box"
      >
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
    </div>
  );
}

export default App;
