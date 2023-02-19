# 圈圈叉叉的遊戲

github.com/limiu82214/GoBasicProject/ooxx

## 規劃

### [ ] 階段一

* 用cmmandline玩
* leveldb
* github.com/c-bata/go-prompt 當作輸入

* [v] 單例board
* [v] board可以設定狀態
* [v] 要輪流下
* [v] 不能重覆下同個位置
* [v] 有一方勝出要立刻顯示
* [v] 勝出後要自動重置狀態
* [v] leveldb存board
* 麻煩：`error`最好一開始就開好
* 麻煩：為每個方法都注入db很麻煩
* 好處：錯誤很明確，直接可以去對應的地方處理
* 好處：當你錯誤引用的時候，感覺蠻明顯的(EX:從service使直接import adapter.out的時候)
* 判斷：LoadBoardPort不應像Board把方法拆出來，因為他並不是主要的業務邏輯，而是把資料存取的地方
* 是否該在db的port也把一個個方法對應還是共用LoadPort就可以了。
* [] 可以同時的多個board
* [] 加入player
* [] 由player輪流玩
* [] 加入三戰兩勝的機制

### [ ] 階段二

* 在mac上用GUI玩
* leveldb
* ebitengine

## [ ] 階段三

* 用任一種可連線方式與同友一起玩

## Version

## Point

* 六角架構
