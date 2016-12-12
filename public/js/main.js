function save() {
  var toSave = $("#save input").val();
  if (toSave != "") {
    $.post("/tasks", {task: toSave}, function (res, err) {
      $('#tasks').append(toSave + "<br>");
    })
  }
}

function textAreaContaining(content, id) {
  var node = document.createElement('textarea');
  node.classList.add('sticky');
  node.id = id;
  node.style.width = '500px';
  node.style.height = '200px';
  var textNode = document.createTextNode(content);

  node.appendChild(textNode);
  return node;
}

function createStickies(contentList) {
  var stickyContainer = $('#stickies')[0];
  contentList.forEach(function (content, index) {
    stickyContainer.appendChild(textAreaContaining(content, index));
  });
}

function allTask() {
  $.get("/tasks", function (res, err) {
    res = [
      'some',
      'thing',
      'foo',
      'bar',
      'hello',
      'world'
    ];
    createStickies(res);
  })
}
window.load = allTask();

