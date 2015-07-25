var api = window.api || {};

api.hello = function(){
  window.location.href = '/hello';
};

api.panic = function(){
  window.location.href = '/panic';
};

api.insta = {};

api.insta.philippines = function(){
  var content = document.getElementById('content');
  var xhr = new XMLHttpRequest();
  xhr.open('GET', '/insta/philippines', false);
  xhr.send();
  if (xhr.status != 200) {
    console.error( xhr.status + ': ' + xhr.statusText );
  } else {
    console.log( xhr.responseText );
    var json = JSON.parse(xhr.responseText);
    for (var i=0; i<json.data.length; i++){
      var item = json.data[i];

      var url = item.Images.standard_resolution.Url;
      var img = document.createElement("img");
      img.src = url;
      content.appendChild(img);
    }
  }
};
