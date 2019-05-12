$(document).ready(function(){
  $(".alert").hide();
  Table_1_Init();
  Table_Abolition_Project_Init();
  RunningProjectTable_Init();
  FinishedProjectTable_Init();
  $.ajaxSetup({
    contentType: "application/x-www-form-urlencoded; charset=utf-8"
  });
  var last_clicked_abolition_project_id;
  var last_clicked_apply_project_id;
  var information_modal_type;
  var last_clicked_running_project_id;
  var last_clicked_running_project_leftover_funds;

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
    $("#u-name").html("执行中")
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
      var file = $('#verify-inviterway-file')[0];
      if(file.outerHTML){
        file.outerHTML = file.outerHTML;
      }else{
        file.value = '';
      }
    })

    //招标信息模态框
    $("#modal-invite-project").on('hide.bs.modal',function() {
      $("#invite-project-file").empty();
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
            url: "http://localhost:8080/project/teacher/project/detail",
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
            $("#verify-project-id").html(flag.data.project.id);
            $("#verify-project-name").html(flag.data.project.name);
            $("#project-fin-funds").val(flag.data.project.fin_funds);
            $("#verify-project-inviteway").html(flag.data.project.invite_way);
          } 
          $("#modal-verify-project").modal("show");
        },

        "click #tmpTableInvite":function(e,value,row,index){
          last_clicked_apply_project_id = row.id;
          if (row.status != "招投标"){
            alert("此阶段不可查看招投标信息！");
            return;
          }
          var flag;
          $.ajax({
            url: "http://localhost:8080/project/teacher/project/invite",
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
            $("#invite-project-id").html(flag.data.id);
            $("#invite-project-name").html(flag.data.name);
            $("#invite-project-begintime").html(flag.data.begin_time);
            $("#invite-project-budget").html(flag.data.funds);
            $("#invite-project-inviteway").html(flag.data.invite_way);
            $("#invite-project-instruction").val(flag.data.instruction);

            var fileName = flag.data.invite_file_name;
            var pre = fileName.split(".")[0];
            var info = "<a href='../static/file/"+row.id+"/"+fileName+"' download='"+fileName+"'>"+pre+"</a>";
            $("#invite-project-file").append(info);
            $("#modal-invite-project").modal("show");
          } else {
            alert(flag.msg);
            return;
          } 
        },

        "click #tmpTableRun":function(e,value,row,index){
          last_clicked_apply_project_id = row.id;
          if (row.status != "招投标"){
            alert("未招标项目不可执行！");
            return;
          }

          $("#modal-verify-project-run").modal("show");
        },
      }
  
      function AddTmpTableFuncAlty(value,row,index){
        return[
          '<button id="tmpTableDelete" type="button" class="btn btn-default">删除</button> &nbsp',
          '<button id="tmpTableVerify" type="button" class="btn btn-default">核定参数</button> &nbsp',
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
              field: 'purpose',
              title: '项目用途'
            },{
              field: 'budget',
              title: '计划预算'
            },{
              field: 'invite_way',
              title: '招标方式'
            },{
              field: 'p_function',
              title: '主要功能'
            },{
              field: 'instruction',
              title: '补充说明'
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
      var purpose = $("input[name='project-purpose']:checked").val();
      var p_function = $("#project-function").val();
      var result = $("#project-expect-result").val();
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
      if (name==""||organization==""||instruction==""||budget==""||p_function==""||result==""||inviteway==undefined||purpose==undefined){
        $("#pwd-warning").html("请将项目信息填写完整！");
        $(".alert").show();
        var timeout=setTimeout(function () {
          $("#pwd-warning").html("");
          $(".alert").hide();
        }, 1500);
        return;
      }
      var flag;
      var project={"name": name, "organization": organization, "instruction": instruction, "budget": parseInt(budget), "inviteway": inviteway,"purpose": purpose,"p_function": p_function,"result":result};
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
      $("#project-function").val("");
      $("#project-expect-result").val("");
      $("input[name='invite-radio']").removeAttr('checked');
      $("input[name='project-purpose']").removeAttr('checked');
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

    function getFileName(o){
      var pos=o.lastIndexOf("\\");
      return o.substring(pos+1);  
    }

    //核定参数确认招标按钮
    $("#btn-modal-verify-project-ok").click(function(){
      var inviteway_instruction = $("#verify-inviteway-instruction").val();
      if (inviteway_instruction==""){
        alert("请填写招标说明！");
        return;
      }
      var id = $("#verify-project-id").html();
      var file = $('#verify-inviterway-file')[0].files[0];
      var filename = getFileName($('#verify-inviterway-file').val());

      var formData = new FormData();
      formData.append("file",file);
      formData.append("id",id);
      formData.append("filename",filename);
      formData.append("instruction",inviteway_instruction)

      var flag; 
      $.ajax({ 
        type : "POST", 
        url : hostip+"project/teacher/project/verify", 
        processData: false,
        contentType: false,
        async: false,
        dataType : "JSON",
        data: formData,
        success: function(data) {
          flag = data;
        },
        error: function (jqXHR) { 
          flag = jqXHR.responseJSON;
        }
      },);
      if (flag.msg == "success"){
        $("#TmpProjectTable").bootstrapTable('refresh');
        $("#modal-verify-project").modal("hide");
        return;
      } else {
        alert(flag.msg);
      }
    });

    $("#btn-apply-change-inviteproject").click(function(){
      $("#apply-change-inviteproject").modal("show");
    });

    //申请修改参数模态框关闭
    $('#apply-change-inviteproject').on('hide.bs.modal',function() {
      $("#apply-change-instruction").val("");
    })

    //申请修改参数
    $("#btn-apply-change-ok").click(function(){
      var instruction = $("#apply-change-instruction").val();
      if (instruction == ""){
        alert("请填写申请说明");
        return;
      }
      var flag; 
      $.ajax({ 
        url : hostip+"project/teacher/project/invite/change", 
        type: "POST",
        async: false,
        dataType : "JSON",
        data: {"id":last_clicked_apply_project_id,"instruction":instruction},
        success: function(data) {
          flag = data;
        },
        error: function (jqXHR) { 
          flag = jqXHR.responseJSON;
        }
      },);
      if (flag.msg == "success"){
        $("#apply-change-inviteproject").modal("hide");
        return;
      } else {
        alert(flag.msg);
      }
    });

    //确认执行项目模态框关闭
    $('#modal-verify-project-run').on('hide.bs.modal',function() {
      $("#project-invite-company").val("");
      $("#project-invite-funds").val("");
    })

    $("#btn-project-run-ok").click(function(){
      var company = $("#project-invite-company").val();
      var funds = $("#project-invite-funds").val();
      if (company==""||funds==""){
        alert("请将信息填写完整!");
        return;
      }
      var r=/^[1-9][0-9]+$/gi;
      if (!r.test(funds)){
        alert("最终资金请输入纯数字！");
        return;
      }
      
      var flag; 
      $.ajax({ 
        url : hostip+"project/teacher/project/run", 
        type: "POST",
        async: false,
        dataType : "JSON",
        data: {"company":company,"funds":funds, "id":last_clicked_apply_project_id},
        success: function(data) {
          flag = data;
        },
        error: function (jqXHR) { 
          flag = jqXHR.responseJSON;
        }
      },);
      if (flag.msg == "success"){
        $("#modal-verify-project-run").modal("hide");
        $("#TmpProjectTable").bootstrapTable('refresh');
        return;
      } else {
        alert(flag.msg);
        return;
      }
    });


    //正在执行的项目
    function RunningProjectTable_Init(){
      queryParams = function (params) {
        var temp = {   
        limit: params.limit,   //页面大小
        offset:params.offset
        };
       return temp;
      };

      //Table中按钮绑定事件
      window.abolitionOperateEvents = {
        "click #running-project-addevent":function(e,value,row,index){
          last_clicked_running_project_id = row.id;
          last_clicked_running_project_leftover_funds = row.leftover_funds;
          $("#modal-running-project-addevent").modal("show");
        },
        "click #running-project-listevent":function(e,value,row,index){
          var flag;
          $.ajax({
            url: "http://localhost:8080/project/teacher/run/eventlist",
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
          if (flag.msg != "success"){
            alert(flag.msg);
            return;
          } 

          for (var i=0;i<flag.data.length;i++){
            var info = "<div class='form-group'>\
            <label style='font-weight:bold;'>项目ID:</label> &nbsp;"+"<label>"+flag.data[i].id+"</label> &nbsp; &nbsp;&nbsp; &nbsp;<label style='font-weight:bold;'>项目名称:</label>  &nbsp;<label>"+flag.data[i].name+"</label>\
            <br><label style='font-weight:bold;'>时间:</label> &nbsp;"+flag.data[i].time+"<br><label style='font-weight:bold;'>使用预算:</label> &nbsp;"+"\
            <label>"+flag.data[i].use_funds+"</label><br>\
            <label style='font-weight:bold;'>使用说明:</label> &nbsp; <label>"+flag.data[i].instruction+"</label>\
            </div> <hr>";
            $("#running-project-event-list").append(info);
          }

          $("#modal-running-project-listevent").modal("show");
        },

        "click #running-project-finish":function(e,value,row,index){
          last_clicked_running_project_id = row.id;
          $("#running-project-total-money").val(row.fin_funds);
          $("#running-project-final-used-money").val(row.fin_funds-row.leftover_funds);
          $("#modal-running-project-finish").modal("show");
        },
      }

      function AddAbolitionTableFuncAlty(value,row,index){
        return[
          '<button id="running-project-addevent" type="button" class="btn btn-default">添加事件</button>',
          '<button id="running-project-listevent" type="button" class="btn btn-default">查看事件</button>',
          '<button id="running-project-finish" type="button" class="btn btn-default">完成项目</button>',
        ].join("")
      };

      $('#RunningProjectTable').bootstrapTable({
        url: 'http://localhost:8080/project/teacher/project/run',         //请求后台的URL（*）
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
          field: 'run_time',
          title: '开始时间',
          sortable: true
        },{
          field: 'company_name',
          title: '中标公司'
        },{
          field: 'fin_funds',
          title: '计划预算'
        },{
          field: 'leftover_funds',
          title: '剩余预算'
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

    $("#btn-refresh-RunningProjectTable").click(function(){
      $("#RunningProjectTable").bootstrapTable('refresh');
    });

    //添加事件模态框关闭
    $('#modal-running-project-addevent').on('hide.bs.modal',function() {
      $("#running-project-use-instruction").val("");
      $("#running-project-use-money").val("");
    })

    $("#running-project-addevent-ok").click(function(){
      instruction = $("#running-project-use-instruction").val();
      useFunds = $("#running-project-use-money").val();
      if (instruction==""||useFunds==""){
        alert("请把信息填写完整！");
        return;
      }
      var r=/^[1-9][0-9]+$/gi;
      if (!r.test(useFunds)){
        alert("最终资金请输入纯数字！");
        return;
      }
      if(useFunds>last_clicked_running_project_leftover_funds){
        alert("使用金额超过剩余预算！");
        return;
      }
      var flag; 
      $.ajax({ 
        url : hostip+"project/teacher/project/run/addevent", 
        type: "POST",
        async: false,
        dataType : "JSON",
        data: {"id":last_clicked_running_project_id,"funds":useFunds, "instruction":instruction},
        success: function(data) {
          flag = data;
        },
        error: function (jqXHR) { 
          flag = jqXHR.responseJSON;
        }
      },);
      if (flag.msg == "success"){
        $("#modal-running-project-addevent").modal("hide");
        $("#RunningProjectTable").bootstrapTable('refresh');
        return;
      } else {
        alert(flag.msg);
        return;
      }
    });

    //查看事件模态框关闭
    $('#modal-running-project-listevent').on('hide.bs.modal',function() {
      $("#running-project-event-list").empty();
    })

    //完成项目
    $("#btn-running-project-finish").click(function(){
      var flag; 
      $.ajax({ 
        url : hostip+"project/teacher/project/run/finish", 
        type: "POST",
        async: false,
        dataType : "JSON",
        data: {"id":last_clicked_running_project_id},
        success: function(data) {
          flag = data;
        },
        error: function (jqXHR) { 
          flag = jqXHR.responseJSON;
        }
      },);
      if (flag.msg == "success"){
        $("#modal-running-project-finish").modal("hide");
        $("#RunningProjectTable").bootstrapTable('refresh');
        return;
      } else {
        alert(flag.msg);
        return;
      }
    });

    //已完成项目
    $("#btn-refresh-FinishedProjectTable").click(function(){
      $("#FinishedProjectTable").bootstrapTable('refresh');
    });

    //正在执行的项目
    function FinishedProjectTable_Init(){
      queryParams = function (params) {
        var temp = {   
        limit: params.limit,   //页面大小
        offset:params.offset
        };
       return temp;
      };

      //Table中按钮绑定事件
      window.abolitionOperateEvents = {
        "click #finished-project-detail":function(e,value,row,index){
        },

      }

      function AddAbolitionTableFuncAlty(value,row,index){
        return[
          '<button id="finished-project-detail" type="button" class="btn btn-default">详情</button>',
          '<button id="finished-project-inviteinfo" type="button" class="btn btn-default">招标信息</button>',
        ].join("")
      };

      $('#FinishedProjectTable').bootstrapTable({
        url: 'http://localhost:8080/project/teacher/project/finished',         //请求后台的URL（*）
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
          field: 'run_time',
          title: '开始时间',
          sortable: true
        },{
          field: 'fin_time',
          title: '结束时间'
        },{
          field: 'fin_funds',
          title: '计划预算'
        },{
          field: 'used_funds',
          title: '使用预算'
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
});