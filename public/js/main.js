function save(content) {
  $.post('/tasks', {task: content}, function (res, err) {
  })
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

function createStickies(contentList, id) {
  var stickyContainer = $('#stickies')[0];
  contentList.forEach(function (content, index = id || index) {
    stickyContainer.prepend(textAreaContaining(content, index));
  });
}

function allTask() {
  $.get('/tasks', function (res, err) {
    res = res.split('<br/>');
    res.length -= 1;
    createStickies(res);
  })
}

function addSticky() {
  $('#add').hide();
  var newBlock = $('#new');
  var listOfStickies = $('.sticky');
  var id = listOfStickies.length != 0 ? Number(listOfStickies.first()[0].id) + 1 : 1;
  var node = textAreaContaining('', id);
  node.classList.add('new');

  node.style.width = '1600px';
  node.style.height = '500px';

  newBlock.append(node);
  $('#save').show();
}

function reset() {
  var newSticky = $('.new').first();
  newSticky.remove();
  $('#add').show();
  $('#save').hide();

}

function saveSticky() {
  var newSticky = $('.new').first();
  var content = newSticky.val();
  createStickies([content], newSticky.id);
  save(content);
  reset();
}

window.load = allTask();

