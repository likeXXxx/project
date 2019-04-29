$(document).ready(function(){
  $(".alert").hide();
  Table_1_Init();
  Table_Abolition_Project_Init();
  $.ajaxSetup({
    contentType: "application/x-www-form-urlencoded; charset=utf-8"
  });
  var last_clicked_abolition_project_id;
  var last_clicked_apply_project_id;
  var information_modal_type;

  var TeacherInfo = new Array(2);
  var hostip = "http://localhost:8080/";
  var flag; 
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
    TeacherInfo[0] = flag.data[0];
    TeacherInfo[1] = flag.data[1];
    $("#userinfo-btn").html(flag.data[0]);
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

  
    //核定参数模态框关闭
    $('#modal-verify-project').on('hide.bs.modal',function() {
      $("#verify-project-id").html("");
      $("#verify-project-name").html("");
      $("#project-fin-funds").val("");
      $("#verify-project-inviteway").html("");
      $("#verify-inviteway-instruction").val("");
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
        "click #tmpTableDelete":function(e,value,row,index){
          if (row.status != "学院审核") {
            alert("已通过审核的项目不能删除！");
            return;
          }
          last_clicked_apply_project_id = row.id;
          information_modal_type = "apply";
          $("#modal-information").modal("show");
        },

        "click #tmpTableVerify":function(e,value,row,index){
          if (row.status != "核定参数"){
            alert("此阶段不可核定参数！");
            return;
          }
          var flag;
              $.ajax({
              url: "http://localhost:8080/project/omanager/project/detail",
              type: "GET",
              async: false,
              dataType : "JSON",
              data: {"id": row.id},
              success: function(data) {
                      flag = data;
                    },
              error: function (jqXHR) { 
                flag = jqXHR.responseJSON;
              }
              },);
              if (flag.msg == "success"){

              } 
          $("#modal-verify-project").modal("show");
        }
      }
  
      function AddTmpTableFuncAlty(value,row,index){
        return[
          '<button id="tmpTableDelete" type="button" class="btn btn-default">删除</button> &nbsp',
          '<button id="tmpTableVerify" type="button" class="btn btn-default">核定参数</button> &nbsp',
          '<button id="tmpTableInvite" type="button" class="btn btn-default">修改参数</button> &nbsp',
          '<button id="tmpTableRun" type="button" class="btn btn-default">确认执行</button> &nbsp'
        ].join("")
      }

        $('#TmpProjectTable').bootstrapTable({
            url: 'http://localhost:8080/project/teacher/project/temp',         //请求后台的URL（*）
            method: 'get',                      //请求方式（*）
            striped: true,                      //是否显示行间隔色
            cache: false,                       //是否使用缓存，默认为true，所以一般情况下需要设置一下这个属性（*）
            pagination: true,                   //是否显示分页（*）
            sortable: true,                     //是否启用排序
            sortOrder: "asc",                   //排序方式
            sidePagination: 'client',           //分页方式：client客户端分页，server服务端分页（*）
            pageNumber: 1,                       //初始化加载第一页，默认第一页
            pageSize: 10,                       //每页的记录行数（*）
            pageList: [10, 15, 20],        //可供选择的每页的行数（*）
            queryParams: queryParams,           //传递参数（*）
            search: true,                       //是否显示表格搜索，此搜索是客户端搜索，不会进服务端，所以，个人感觉意义不大
            contentType: "application/x-www-form-urlencoded",
            strictSearch: true,
            showColumns: true,                  //是否显示所有的列
            showRefresh: true,                  //是否显示刷新按钮
            clickToSelect: true,                //是否启用点击选中行
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
              title: '创建时间',
              sortable: true
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
              field: 'status',
              title: '项目状态'
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
    };

    $("#btn-refresh-tmptable").click(function(){
      $("#TmpProjectTable").bootstrapTable('refresh');
    });

    $("#btn-new-project").click(function(){
      $("#project-owner").html(TeacherInfo[0]);
      $("#project-organization").html(TeacherInfo[1]);
    });

    $("#btn-newproject-ok").click(function(){
      var name = $("#project-name").val();
      var organization = $("#project-organization").html();
      var instruction = $("#project-instruction").val();
      var budget = $("#project-budget").val();
      var inviteway = $("input[name='invite-radio']:checked").val(); 
      if (inviteway == "其他方式") {
        inviteway = $("#project-invite-way-other").val();
      }
      var r=/^[1-9][0-9]+$/gi;
      if (!r.test(budget)){
        $("#pwd-warning").html("预算请输入纯数字！");
        $(".alert").show();
        var timeout=setTimeout(function () {
          $("#pwd-warning").html("");
          $(".alert").hide();
        }, 1500);
        return;
      }
      if (name==""||organization==""||instruction==""||budget==""||inviteway==undefined){
        $("#pwd-warning").html("请将项目信息填写完整！");
        $(".alert").show();
        var timeout=setTimeout(function () {
          $("#pwd-warning").html("");
          $(".alert").hide();
        }, 1500);
        return;
      }
      var flag;
      var project={"name": name, "organization": organization, "instruction": instruction, "budget": parseInt(budget), "inviteway": inviteway};
	    $.ajax({
		    url: hostip+"project/teacher/project",
		    type: "POST",
        async: false,
        dataType: "json",
        contentType: "application/json",
		    data: JSON.stringify(project),
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
        $("#modal-new-project").modal('hide');
        $("#TmpProjectTable").bootstrapTable('refresh');
        return;
      }
    });

    $('#modal-new-project').on('hide.bs.modal',
    function() {
      $("#project-name").val("");
      $("#project-instruction").val("");
      $("#project-budget").val("");
      $("#project-invite-way-other").val("");
      $("input[name='invite-radio']").removeAttr('checked');
    });

    

    function Table_Abolition_Project_Init(){
      queryParams = function (params) {
        var temp = {   
        limit: params.limit,   //页面大小
        offset:params.offset
        };
       return temp;
      };

      //Table中按钮绑定事件
      window.abolitionOperateEvents = {
        "click #abolition-project-delete":function(e,value,row,index){
          last_clicked_abolition_project_id = row.id;
          information_modal_type = "abolition"
          $("#modal-information").modal("show");
        }
      }

      function AddAbolitionTableFuncAlty(value,row,index){
        return[
          '<button id="abolition-project-delete" type="button" class="btn btn-default">删除</button>',
        ].join("")
      };

      $('#abolition-project-table').bootstrapTable({
        url: 'http://localhost:8080/project/teacher/project/abolition',         //请求后台的URL（*）
        method: 'get',                      //请求方式（*）
        striped: true,                      //是否显示行间隔色
        cache: false,                       //是否使用缓存，默认为true，所以一般情况下需要设置一下这个属性（*）
        pagination: true,                   //是否显示分页（*）
        sortable: true,                     //是否启用排序
        sortOrder: "asc",                   //排序方式
        sidePagination: 'client',           //分页方式：client客户端分页，server服务端分页（*）
        pageNumber: 1,                       //初始化加载第一页，默认第一页
        pageSize: 10,                       //每页的记录行数（*）
        pageList: [10, 15, 20],        //可供选择的每页的行数（*）
        queryParams: queryParams,           //传递参数（*）
        search: true,                       //是否显示表格搜索，此搜索是客户端搜索，不会进服务端，所以，个人感觉意义不大
        contentType: "application/x-www-form-urlencoded",
        strictSearch: true,
        showColumns: true,                  //是否显示所有的列
        showRefresh: true,                  //是否显示刷新按钮
        clickToSelect: true,                //是否启用点击选中行
        uniqueId: "no",                     //每一行的唯一标识，一般为主键列
        showToggle: true,                    //是否显示详细视图和列表视图的切换按钮
        cardView: false,                    //是否显示详细视图
        detailView: false,                   //是否显示父子表
        columns: [
        {
          field: 'id',
          title: 'ID'
        }, {
          field: 'name',
          title: '名称'
        }, {
          field: 'create_time',
          title: '创建时间',
          sortable: true
        },{
          field: 'abolition_organization',
          title: '终止机构'
        },{
          field: 'abolition_instruction',
          title: '终止说明'
        },{
          field: 'operator',
          title: '操作人'
        },{
          field: 'operator_tel',
          title: '联系方式'
        },{
          field: 'useroperator',
          title: '操作',
          events: abolitionOperateEvents,
          formatter: AddAbolitionTableFuncAlty
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

    //删除模态框点击确认
    $("#btn-information-modal-ok").click(function(){
      if(information_modal_type == "abolition"){
        var flag; 
        $.ajax({
          url: hostip+"project/teacher/project/abolition",
          type: "delete",
          async: false,
          dataType: "JSON",
          contentType: "application/json",
		      data:JSON.stringify({"id": last_clicked_abolition_project_id}),
          success: function(data) {
            flag = data;
          },
          error: function (jqXHR) { 
            flag = jqXHR.responseJSON;
          }
        },);
        if (flag.msg == "success"){
          $("#abolition-project-table").bootstrapTable('refresh');
          $("#modal-information").modal("hide");
          return;
        } else {
          alert(flag.msg);
        }
      } else if (information_modal_type == "apply"){
        var flag; 
        $.ajax({
          url: hostip+"project/teacher/project",
          type: "DELETE",
          async: false,
          contentType:"application/json",
          dataType : "JSON",
		      data:JSON.stringify({"id": last_clicked_apply_project_id}),
          success: function(data) {
            flag = data;
          },
          error: function (jqXHR) { 
            flag = jqXHR.responseJSON;
          }
        },);
        if (flag.msg == "success"){
          $("#TmpProjectTable").bootstrapTable('refresh');
          $("#modal-information").modal("hide");
          return;
        } else {
          alert(flag.msg);
        }
      }
    });

    //核定参数确认招标按钮
    $("#btn-modal-verify-project-ok").click(function(){

    });


});