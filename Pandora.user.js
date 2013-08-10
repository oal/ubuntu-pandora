// ==UserScript==
// @name        Pandora
// @namespace   pandora
// @description Pandora
// @include     http://www.pandora.com/station/play/*
// @version     1
// ==/UserScript==

// Change the line below to point to the Go server:
var host = 'http://192.168.0.101:10002';

var lastSong = '';
setInterval(function() {
    var thisSong = $('.songTitle').text();
    
    if(thisSong != lastSong) {
        $.post(host, {
            'song': thisSong,
            'album': $('.albumTitle').text(),
            'artist': $('.artistSummary').text()
        })
        lastSong = thisSong;
    };
}, 1000);
