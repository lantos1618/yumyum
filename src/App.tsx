
import { useEffect, useState } from "react";
import { YumYumServiceClient } from "./generated/proto/YumyumServiceClientPb";
import { Emoji, EmojiReaction } from "./generated/proto/yumyum_pb";

const client = new YumYumServiceClient("http://localhost:8080", null, null);


function App() {
  const [emojiReaction, setEmojiReaction] = useState(EmojiReaction.LOVE);

  useEffect(() => {
    const request = new Emoji();
    request.setReaction(emojiReaction);

    client.emojiChat(request, {}, (err, response) => {
      if (err) {
        console.log(`Unexpected error for emojiChat: code = ${err.code}, message = "${err.message}"`);
      } else {
        setEmojiReaction(response.getReaction());
      }
    });
  }, [emojiReaction]);

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
