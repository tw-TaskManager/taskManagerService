function save(){
    var toSave = $("#save input").val()
    $.post("/save",{task:toSave},function(res,err){
        alert(res);
    })
}
