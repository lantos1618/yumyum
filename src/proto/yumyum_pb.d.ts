import * as jspb from 'google-protobuf'



export class Emoji extends jspb.Message {
  getReaction(): EmojiReaction;
  setReaction(value: EmojiReaction): Emoji;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Emoji.AsObject;
  static toObject(includeInstance: boolean, msg: Emoji): Emoji.AsObject;
  static serializeBinaryToWriter(message: Emoji, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Emoji;
  static deserializeBinaryFromReader(message: Emoji, reader: jspb.BinaryReader): Emoji;
}

export namespace Emoji {
  export type AsObject = {
    reaction: EmojiReaction,
  }
}

export class Empty extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Empty.AsObject;
  static toObject(includeInstance: boolean, msg: Empty): Empty.AsObject;
  static serializeBinaryToWriter(message: Empty, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Empty;
  static deserializeBinaryFromReader(message: Empty, reader: jspb.BinaryReader): Empty;
}

export namespace Empty {
  export type AsObject = {
  }
}

export enum EmojiReaction { 
  UNKNOWN = 0,
  LIKE = 1,
  LOVE = 2,
  HAHA = 3,
  WOW = 4,
  SAD = 5,
  ANGRY = 6,
}
