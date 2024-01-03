package libs

func RemoveElm[T any](sl []T, idx int) []T {
	if (idx + 1) > len(sl) {
		return sl
	}
	if (idx + 1) == len(sl) {
		return sl[:idx]
	}

	// idx以降を、idx+1以降で上書きする（slの末尾には変更がされない）
	changed := copy(sl[idx:], sl[idx+1:])

	// slの末尾を削除する
	return sl[:idx+changed]
}
