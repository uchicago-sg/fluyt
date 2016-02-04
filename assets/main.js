var addMore = function(done) {
  var xhr = new XMLHttpRequest;
  xhr.open("POST", "/", true);
  xhr.setRequestHeader("Content-type", "application/www-form-urlencoded");
  xhr.setRequestHeader("Accept", "text/javascript");
  xhr.onreadystatechange = function() {
    if (xhr.readyState == 4) {
      done();
    }
  };
  xhr.send(null);
}

var render = function(query) {
  var body = document.body;
  while (body.hasChildNodes())
    body.removeChild(body.firstChild);

  var input = document.createElement("input");
  input.value = query;
  input.className = "form-control";
  input.onchange = function() {
    render(input.value);
  };
  input.select();
  input.focus();
  body.appendChild(input);

  var button = document.createElement("button");
  button.innerText = "Add 1,000 Listings";
  button.className = "btn btn-default";
  button.onclick = function() {
    addMore(function() {
      button.disabled = false;
    });
    button.disabled = true;
  }
  body.appendChild(button);

  var xhr = new XMLHttpRequest;
  xhr.open("GET", "/?q=" + query, true);
  xhr.setRequestHeader("Accept", "text/javascript");
  xhr.onreadystatechange = function() {
    if (xhr.readyState == 4) {
      var listings = JSON.parse(xhr.responseText).listings;

      var table = document.createElement("table"); table.className = "table table-striped";
      body.appendChild(table);

      var tr = document.createElement("tr");
      var th = document.createElement("th"); tr.appendChild(th);
      th.innerText = listings.length;
      tr.appendChild(th);
      table.appendChild(tr);
      
      var tbody = document.createElement("tbody"); table.appendChild(tbody);

      
      for (var i = 0; i < listings.length; i++) {
        var tr = document.createElement("tr");
        var td = document.createElement("td"); tr.appendChild(td);
        td.innerText = listings[i].key;
        tr.appendChild(td);
        tbody.appendChild(tr);
      }
    }
  };
  xhr.send(null);
  
};

window.onload = function() {
  var qs = window.location.search.substring(1).split("&");
  var query = {};
  for (var i = 0; i < qs.length; i++) {
    var bits = qs[i].split("=");
    query[decodeURIComponent(bits[0])] = decodeURIComponent(bits[1]);
  }

  render(query["q"] || "");
};