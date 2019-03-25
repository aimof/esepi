package main

const startTemplate = `<html lang="ja">
    <header>
        <meta charset="UTF-8">
        <meta title="ESEPI適性試験">
        <link rel="stylesheet" href="bootstrap.min.css">
    </header>
    <body>
        <div class="container">
            <div class="row">
                <div class="col-8 col-push-2">
                    <h1>ESEPI適性試験 ver0.0.0</h1>
                </div>
            </div>
            <div class="row">
                <div class="col-8 col-push-2">
                    <form action="question">
                        <input type="hidden" name="number" value="0">
                        <div class="form-group">
                            <label>名前またはIDを入力してください。</label>
                            <input type="text" name="id">
                        </div>
                        <div class="form-group">
                            <label>試験開始ボタンを押すと計測が始まります。</label>
                            <button class="btn btn-primary">試験開始</button>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </body>
</html>`
