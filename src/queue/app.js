var express = require('express');
var app = express();
var bodyParser = require('body-parser');
var fs = require('fs');
app.use(bodyParser.urlencoded({ extended: true}));
app.use(bodyParser.json());

app.get('/', function(req, res) {
	res.sendFile(__dirname + '/view/index.html');
});

app.post('/', function(req, res) {
    const blob = req.body.blob;
    var img = fs.readFileSync(__dirname + '/view/ok.png');
    res.writeHead(200, {'Content-Type': 'image/gif'});
    res.end(img, 'binary');
});

app.listen(8080, function() {
	    console.log("Listening on 8080...");
});
