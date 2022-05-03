import React from "react";
import { Player, Audio, DefaultUi } from "@vime/react";

export const PlayerComponent = () => (
  <Player theme="light" controls>
    <Audio>
        <source src="file:///home/hitanshu/Music/sample-1.mp3" type="audio/mp3" />
    </Audio>
    <DefaultUi />
  </Player>
);
