$(document).ready(function(){
    $(".alert").hide();
    Table_1_Init();
    $.ajaxSetup({
      contentType: "application/x-www-form-urlencoded; charset=utf-8"
    });
  
    var TeacherInfo = new Array(2);
    var flag;
    var hostip = "http://localhost:8080/" 
    $.ajax({
      url: hostip+"project/omanager/getinfo",
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
      TeacherInfo[0] = flag.data[0];
      TeacherInfo[1] = flag.data[1];
      $("#userinfo-btn").html(flag.data[0]);
    } 
  
    $("#li1").click(function(){
      $("#u-name").html("学院申报项目")
    });
    
    $("#li2").click(function(){
      $("#u-name").html("学院执行项目")
    });
    
    $("#li3").click(function(){
      $("#u-name").html("学院完成项目")
    });
  
    $("#item-logout").click(function(){
      $.ajax({
        url: hostip+"project/login/logout",
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
            url: "http://localhost:8080/project/omanager/pwd",
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
  
        window.applyOperateEvents = {
            "click #applyTableDetail":function(e,value,row,index){
                $("#modal-apply-project-detail").modal('show');
            },

            "click #applyTableAuditing":function(e,value,row,index){
                $("#modal-apply-project-Auditing").modal('show');
            }
        }
    
        function AddApplyTableFuncAlty(value,row,index){
          return[
            '<button id="applyTableDetail" type="button" class="btn btn-default">详情</button> &nbsp',
            '<button id="applyTableAuditing" type="button" class="btn btn-default">审核</button> &nbsp',
          ].join("")
        }
  
          $('#ApplyProjectTable').bootstrapTable({
              url: 'http://localhost:8080/project/omanager/project/apply',         //请求后台的URL（*）
              method: 'get',                      //请求方式（*）
              striped: true,                      //是否显示行间隔色
              cache: false,                       //是否使用缓存，默认为true，所以一般情况下需要设置一下这个属性（*）
              pagination: true,                   //是否显示分页（*）
              // sortable: true,                     //是否启用排序
              // sortOrder: "asc",                   //排序方式
              sidePagination: "client",           //分页方式：client客户端分页，server服务端分页（*）
              pageNumber: 1,                       //初始化加载第一页，默认第一页
              pageSize: 5,                       //每页的记录行数（*）
              pageList: [5, 10, 25, 50, 100],        //可供选择的每页的行数（*）
              queryParams: queryParams,           //传递参数（*）
              search: true,                       //是否显示表格搜索，此搜索是客户端搜索，不会进服务端，所以，个人感觉意义不大
              contentType: "application/x-www-form-urlencoded",
              strictSearch: true,
              showColumns: true,                  //是否显示所有的列
              showRefresh: true,                  //是否显示刷新按钮
              clickToSelect: true,                //是否启用点击选中行
              height: 700,                        //行高，如果没有设置height属性，表格自动根据记录条数觉得表格高度
              uniqueId: "no",                     //每一行的唯一标识，一般为主键列
              // showToggle: true,                    //是否显示详细视图和列表视图的切换按钮
              // cardView: false,                    //是否显示详细视图
              // detailView: false,                   //是否显示父子表
              columns: [
              {
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
                field: 'teacher_name',
                title: '负责教师'
              },{
                  field: 'teacher_tel',
                  title: '联系方式'
              },{
                field: 'operator',
                title: '操作',
                events: applyOperateEvents,
                formatter: AddApplyTableFuncAlty,
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
      };
  
      $("#btn-refresh-applytable").click(function(){
        $("#ApplyProjectTable").bootstrapTable('refresh');
      });
  
  });