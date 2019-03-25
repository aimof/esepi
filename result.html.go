package main

const resultTemplate = `<html lang="ja">
    <header>
        <meta charset="UTF-8">
        <meta title="ESEPI適性試験">
        <link rel="stylesheet" href="bootstrap.min.css">
    </header>
    <body>
        <div class="container">
            <div class="row">
                <div class="col-8 col-push-2">
                    <h1>ESEPI適性試験 Ver0.0.0 結果</h1>
                </div>
            </div>
            <div class="row">
                <div class="col-8 col-push-2">
                    <p>この画面のスクリーンショットを撮影してください。</p>
                </div>
                <div class="col-8 col-push-2">
                    <p>受験者：　{{.ID}}</p>
                    <p>正答率：　{{.Accuracy}}%</p>
                    <p>所要時間：　{{.TimeSpent}}秒</p>
                </div>
                <div class="col-8 col-push-2">
                    <a href="https://twitter.com/intent/tweet?button_hashtag=ESEPI&text=適性試験ESEPIであなたは、{{.Accuracy}}％、{{.TimeSpent}}秒の結果を残しました。&url=https%3A%2F%2Fgithub.com%2Faimof%2Fesepi&ref_src=twsrc%5Etfw" class="twitter-hashtag-button" data-show-count="false">#ESEPI の結果をツイート</a><script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>
                </div>
                <div class="col-8 col-push-2">
                    <form action="end">
                        <label class="form-group">スクリーンショット撮影が終わったら終了ボタンを押してください。</label>
                        <label class="form-group"><b>スコアは記録されません。必ずスクリーンショットを撮影してください。</b></label>
                        <button class="btn btn-danger">終了</button>
                    </form>
                </div>
            </div>
        </div>
    </body>
</html>`
