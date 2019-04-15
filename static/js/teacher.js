$(document).ready(function(){
  $(".alert").hide();
  Table_1_Init();

  var flag;
  var hostip = "http://localhost:8080/" 
  $.ajax({
    url: hostip+"project/teacher/getinfo",
    type: "GET",
    async: false,
    success: function(data) {
      flag = data;
      },
    error: function (jqXHR) { 
      flag = jqXHR.responseJSON;
    }
  },);
  if (flag.msg == "success"){
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

  $("#item-logout").click(function(){
    $.ajax({
      url: hostip+"project/teacher/logout",
      type: "POST",
      async: false,
    },);
    window.location.replace(hostip+"/project/login");
  });
  
  $("#btn-resetpwd-ok").click(function(){
    var old = $("#pwd-old").val();
    var newPwd = $("#pwd-new").val();
    var rePwd = $("#pwd-renew").val();
    if (newPwd != rePwd){
      $("#pwd-warning").html("请确认两次输入的新密码相同！");
      $(".alert").show();
      var timeout=setTimeout(function () {
        $("#pwd-warning").html("");
        $(".alert").hide();
      }, 1500);
      return;
    }
    
    var flag;
	  $.ajax({
		  url: "http://localhost:8080/project/teacher/pwd",
		  type: "PUT",
		  async: false,
      dataType : "JSON",
		  data: {"old": old, "new": newPwd},
      success: function(data) {
			  flag = data;
		    },
      error: function (jqXHR) { 
        flag = jqXHR.responseJSON;
      }
    },);
    if (flag.msg != "success"){
      $("#pwd-warning").html(flag.msg);
      $(".alert").show();
      var timeout=setTimeout(function () {
        $("#pwd-warning").html("");
        $(".alert").hide();
      }, 1500);
      return;
    } else {
      window.location.replace(hostip+"/project/login");
    }
  });

  $('#myModal').on('hide.bs.modal',
    function() {
      $("#pwd-old").val("");
      $("#pwd-new").val("");
      $("#pwd-renew").val("");
    })

    function Table_1_Init() {
      //得到查询的参数
      queryParams = function (params) {
        var temp = {   //这里的键的名字和控制器的变量名必须一直，这边改动，控制器也需要改成一样的
        limit: params.limit,   //页面大小
        offset:params.offset
        };
       return temp;
      }

      window.tempOperateEvents = {

      }
  
      function AddTmpTableFuncAlty(value,row,index){
        return[
          '<button id="tmpTableDelete" type="button" class="btn btn-default">删除</button> &nbsp',
          '<button id="tmpTableVerify" type="button" class="btn btn-default">确认参数</button> &nbsp',
          '<button id="tmpTableInvite" type="button" class="btn btn-default">招投标</button> &nbsp',
          '<button id="tmpTableRun" type="button" class="btn btn-default">确认执行</button> &nbsp'
        ].join("")
      }

        $('#TmpProjectTable').bootstrapTable({
            url: 'http://localhost:8080/project/teacher/project/temp',         //请求后台的URL（*）
            method: 'get',                      //请求方式（*）
            striped: true,                      //是否显示行间隔色
            cache: false,                       //是否使用缓存，默认为true，所以一般情况下需要设置一下这个属性（*）
            pagination: true,                   //是否显示分页（*）
            // sortable: false,                     //是否启用排序
            // sortOrder: "asc",                   //排序方式
            sidePagination: "server",           //分页方式：client客户端分页，server服务端分页（*）
            pageNumber: 1,                       //初始化加载第一页，默认第一页
            pageSize: 5,                       //每页的记录行数（*）
            pageList: [10, 25, 50, 100],        //可供选择的每页的行数（*）
            queryParams: queryParams,           //传递参数（*）
            search: true,                       //是否显示表格搜索，此搜索是客户端搜索，不会进服务端，所以，个人感觉意义不大
            contentType: "application/x-www-form-urlencoded",
            strictSearch: true,
            showColumns: true,                  //是否显示所有的列
            showRefresh: true,                  //是否显示刷新按钮
            // clickToSelect: true,                //是否启用点击选中行
            height: 700,                        //行高，如果没有设置height属性，表格自动根据记录条数觉得表格高度
            uniqueId: "no",                     //每一行的唯一标识，一般为主键列
            // showToggle: true,                    //是否显示详细视图和列表视图的切换按钮
            // cardView: false,                    //是否显示详细视图
            // detailView: false,                   //是否显示父子表
            columns: [
            {
              field: 'id',
              title: 'ID'
            }, {
              field: 'name',
              title: '名称'
            }, {
              field: 'create_time',
              title: '创建时间'
            },{
              field: 'budget',
              title: '计划预算'
            },{
              field: 'invite_way',
              title: '招标方式'
            },{
              field: 'instruction',
              title: '项目说明'
            },{
              field: 'operator',
              title: '操作',
              events: tempOperateEvents,
              formatter: AddTmpTableFuncAlty,
            }
            ],
            rowStyle: function (row, index) {
                var classesArr = ['success', 'info'];
                var strclass = "";
                if (index % 2 === 0) {//偶数行
                    strclass = classesArr[0];
                } else {//奇数行
                    strclass = classesArr[1];
                }
                return { classes: strclass };
            },//隔行变色

            
        });
    }
});