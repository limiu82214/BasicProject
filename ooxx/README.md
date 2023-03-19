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
* 選擇：我可以在main直接做show board，也可以把這個方法丟到player裡面，然後再show board。
    * 把show board在main呼叫，算不算跳過player↓
        * player是所有呼叫的入口點嗎？否
            * 邏輯是否會脫離範圍？否，僅單純呼叫
            * player做下棋這件事情才正確？一半，看user透過player棋，還是user直接下棋
        * 結：我希望user是透過player進行遊戲，而board屬於一個工具類型
* ooxx當前記錄`ooxx0.1`：
    board：
        可以下棋與判斷輸贏等基本功能。
        輸入：
            prompt
            player
        輸出：leveldb
    player
        透過board的功能show棋盤。
        輸入： prompt
        輸出： board
* [v] 加入player
* [v] 要先設定玩家nick name才能玩
* 這是一個插入的需求，他並不符合player本身的邏輯，但是符合業務需求，所以加在usecase
* [v] 導入CI
    ooxx不能進行push。
    開發分支要合併到ooxx必需發起Pull Request。
    開發分支必需通過golangci-lint才能merge進ooxx。
* [v] 整理，抽離共用，整理成固定的規則
    Err
        統一使用 github.com/pkg/errors
        一致使用errors.Wrap與errors.New (重覆Wrap會成加記憶體等成本，但是書寫時可以避免需要判斷上下文有沒有Wrap過的問題)
        使用nerror.PrettyError來印出error鏈與trace path
    命名規劃
        為了使引用變的容易，取捨後決定讓資料夾名稱有重覆
            所有會被引用的package依照六角結構命名
            application New+完整名稱
            其他只用New
    Wire
        使用Wire注入
            將DB，UseCase做成Set
    State
        將State分離出來(因為目前player與board都有用到state)
* [] 重新拆分顆粒度
    API 負責處理所有對外的API處理
    User 負責處理所有的玩家資訊
    Game 負責處理所有的遊戲資訊
* [] 由player輪流玩
* [] 加入三戰兩勝的機制
* [] 可以同時的多個board

### [ ] 階段二

* 在mac上用GUI玩
* leveldb
* ebitengine

## [ ] 階段三

* 用任一種可連線方式與同友一起玩

## Version

## Point

* 六角架構

## 額外原則
* 如果兩個組件都是自己的，要實做in與out，提供功能那方把資料給出可用性高的格式，使用功能那方再把資料轉換成自己需要的格式。
* 在每個地方增加 errInHere 給error wrap使用
* Cmd 使用new建構，並要實作IsValid確保是由New產生的物件，在service中必需要使用IsValid。
