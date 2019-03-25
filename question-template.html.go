package main

const questionTemplate = `<html lang="ja">
    <header>
        <meta charset="UTF-8">
        <meta title="ESEPI適性試験">
        <link rel="stylesheet" href="bootstrap.min.css">
    </header>
    <body>
        <div class="container">
            <div class="row">
                <div class="col-8 col-push-2">
                    <h1>第{{.Number}}問</h1>
                </div>
            </div>
            <div class="row">
                <div class="col-8 col-push-2">
                    <p>{{.Question}}</p>
                    <p>{{.Resources}}</p>
                    <form action="question" name="{{.Number}}">
                        <input type="hidden" name="number" value="{{.Number}}">
                        <div class="form-group">
                            <label><h2>回答欄</h2></label>
                            <input type="text" class="form-control" name="answer">
                        </div>
                        <div class="form-group"><button type="submit" class="btn btn-primary">送信</button></div>
                    </form>
                </div>
            </div>
        </div>
    </body>
</html>`
