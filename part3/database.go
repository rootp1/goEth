func (b *Block) Serialize() []byte {
    var result bytes.Buffer
    encoder := gob.NewEncoder(&result)
    err := encoder.Encode(b)
    return result.Bytes()
}
