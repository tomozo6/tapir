package mypkg

/*
タグ包含判定関数

	概要:
	    検索する側のリスト内のタグが、検索される側のリストに
		全て含まれている場合trueを返します。

    引数:
        iTags: 検索する側のタグリスト
		jTags: 検索される側のタグリスト

     例(true):
        iTags: map[project:a env:prod]
		jTags: map[project:a env:prod name:abcde]

     例(false):
        iTags: map[project:b env:prod]
		jTags: map[project:a env:prod name:abcde]

     例(false):
        iTags: map[project:a env:prod role:web]
		jTags: map[project:a env:prod name:abcde]
*/

func tagContains(iTags map[string]string, jTags map[string]string) bool {
	var match_flg bool
	for ii, iv := range iTags {
		match_flg = false
		for ji, jv := range jTags {
			if ii == ji {
				if iv == jv {
					match_flg = true
					break
				}
			}
		}
		if match_flg == false {
			return false
		}
	}
	return true
}
