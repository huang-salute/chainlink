package resolver

import "github.com/smartcontractkit/chainlink/core/services/keystore/keys/ocrkey"

type OCRKey struct {
	configPublicKey       ocrkey.ConfigPublicKey
	offChainPublicKey     ocrkey.OffChainPublicKey
	onChainSigningAddress ocrkey.OnChainSigningAddress
}

func (k OCRKey) ConfigPublicKey() string {
	return k.configPublicKey.String()
}

func (k OCRKey) OffChainPublicKey() string {
	return k.offChainPublicKey.String()
}

func (k OCRKey) OnChainSigningAddress() string {
	return k.onChainSigningAddress.String()
}

type OCRKeysPayloadResolver struct {
	keys []ocrkey.KeyV2
}

func NewOCRKeysPayloadResolver(keys []ocrkey.KeyV2) *OCRKeysPayloadResolver {
	return &OCRKeysPayloadResolver{keys}
}

func (r *OCRKeysPayloadResolver) Keys() []OCRKey {
	viewKeys := []OCRKey{}
	for _, k := range r.keys {
		viewKeys = append(viewKeys, OCRKey{
			configPublicKey:       k.PublicKeyConfig(),
			offChainPublicKey:     k.OffChainSigning.PublicKey(),
			onChainSigningAddress: k.OnChainSigning.Address(),
		})
	}
	return viewKeys
}
