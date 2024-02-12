// Original file: ../proto/paste.proto


export interface PasteRequest {
  'content'?: (string);
  'expiration'?: (string);
  'encrypted'?: (boolean);
}

export interface PasteRequest__Output {
  'content': (string);
  'expiration': (string);
  'encrypted': (boolean);
}
