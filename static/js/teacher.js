$(document).ready(function(){
  var flag;
  var hostip = "http://localhost:8080/" 
  $.ajax({
    url: hostip+"project/teacher/getinfo",
    type: "GET",
    async: false,
    dataType : "JSON",
    data: {"type": type, "num": num, "password": password},
    success: function(data) {
      flag = data;
      },
    error: function (jqXHR) { 
      flag = jqXHR.responseJSON;
    }
  },);
  if (flag.msg == "success"){
    alert("xx");
    $("#userinfo-btn").html(flag.data);
  } 

  $("#li1").click(function(){
    $("#u-name").html("暂存项目")
  });
  
  $("#li2").click(function(){
    $("#u-name").html("申报中")
  });
  
  $("#li3").click(function(){
    $("#u-name").html("已完成项目")
  });
  
});