
import { useEffect, useState } from "react";
import { yumyum } from "./proto/yumyum";
import { grpc } from "@improbable-eng/grpc-web";



const client = new yumyum.YumYumServiceClient("http://localhost:8080", createCre

function App() {



  return (
    <div className="App">
      <header className="App-header">
        <p>
          Emoji Reaction: {emojiReaction}
        </p>
        <button onClick={() => setEmojiReaction(EmojiReaction.LOVE)}>Send Love Reaction</button>
      </header>
    </div>
  );
}

export default App;
