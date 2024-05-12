package locku

type LockEntry struct {
	RequestID string //
	Type      string
	Keys      []string
}

func (le *LockEntry) DeepCopy() LockEntry {
	return LockEntry{
		RequestID: le.RequestID,
		Type:      le.Type,
		Keys:      append(make([]string, 0, len(le.Keys)), le.Keys...),
	}
}
