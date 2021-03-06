$(document).ready(function(){
    $(".alert").hide();
    Table_1_Init();
    Table_ApplyChangeInviteProjectTable_Init();
    $.ajaxSetup({
      contentType: "application/x-www-form-urlencoded; charset=utf-8"
    });
  
    var last_clicked_change_project_id;
    var flag;
    var hostip = "http://localhost:8080/" 
    $.ajax({
      url: hostip+"project/imanager/getinfo",
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
            url: "http://localhost:8080/project/imanager/pwd",
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
  
    //修改密码模态框
    $('#myModal').on('hide.bs.modal',
      function() {
        $("#pwd-old").val("");
        $("#pwd-new").val("");
        $("#pwd-renew").val("");
      });

    //项目详情模态框
    $('#modal-apply-project-detail').on('hide.bs.modal',
      function(){
        $("#project-detail-id").html("");
        $("#project-detail-name").html("");
        $("#project-detail-organization").html("");
        $("#project-detail-budget").html("");
        $("#project-detail-createtime").html("");
        $("#project-detail-inviteway").html("");
        $("#project-detail-instruction").html("");
        $("#teacher-detail-name").html("");
        $("#teacher-detail-id").html("");
        $("#teacher-detail-sex").html("");
        $("#teacher-detail-organization").html("");
        $("#teacher-detail-pt").html("");
        $("#teacher-detail-tel").html("");
        $("#organization-verify-instruction").val("");
        $("#project-detail-purpose").html("");
        $("#project-detail-function").html("");
        $("#project-detail-result").html("");
      });

      var last_clicked_apply_project_id;
      var last_clicked_apply_project_status;
      var last_clicked_fin_project_id;

      //审核项目模态框
      $('#modal-apply-project-Auditing').on('hide.bs.modal',
      function() {
        $("#verify-instruction").val("");
        $("#master-name").val("");
        cleanMasterArray();
      });

      $('#modal-final-auditing').on('hide.bs.modal',
      function() {
        $("#fin-verify-instruction").val("");
        $("#fin-verify-funds").val("");
        $("#fin-modal-body").empty();
      });

      //初审通过
      $('#auditing-pass').click(function(){
        var verifyInstruction = $("#verify-instruction").val();
        var masterInfo = $("#master-name").val();
        if (verifyInstruction == "") {
          alert("审核说明不能为空！");
          return;
        }
        if (masterInfo == ""||(masterInfo.split(",")).length<3) {
          alert("请指定至少三个审核的专家！");
          return;
        }

        var flag;
        $.ajax({
            url: "http://localhost:8080/project/imanager/project/pass",
            type: "POST",
            async: false,
            dataType : "JSON",
            data: {"instruction": verifyInstruction,"id":last_clicked_apply_project_id,"master":masterInfo},
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
        } else {
          $("#ApplyProjectTable").bootstrapTable('refresh');
          $("#modal-apply-project-Auditing").modal('hide');
          return;
        }
      });

      //初审不过
      $('#auditing-fail').click(function(){
        var verifyInstruction = $("#verify-instruction").val();
        if (verifyInstruction == "") {
          alert("审核说明不能为空！");
          return;
        }

        var flag;
        $.ajax({
            url: "http://localhost:8080/project/imanager/project/fail",
            type: "POST",
            async: false,
            dataType : "JSON",
            data: {"instruction": verifyInstruction, "id":last_clicked_apply_project_id},
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
        } else {
          $("#ApplyProjectTable").bootstrapTable('refresh');
          $("#modal-apply-project-Auditing").modal('hide');
          return;
        }
      });
  
      //审核项目table
      function Table_1_Init() {
        //得到查询的参数
        queryParams = function (params) {
          var temp = {   //这里的键的名字和控制器的变量名必须一直，这边改动，控制器也需要改成一样的
          limit: params.limit,   //页面大小
          offset:params.offset
          };
         return temp;
        }
  
        //Table中按钮绑定事件
        window.applyOperateEvents = {
            "click #applyTableDetail":function(e,value,row,index){
              last_clicked_apply_project_id = row.id;
              last_clicked_apply_project_status = row.status;
              var flag;
              $.ajax({
              url: hostip+"project/imanager/project/detail",
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
                $("#project-detail-id").html(flag.data.project.id);
                $("#project-detail-name").html(flag.data.project.name);
                $("#project-detail-organization").html(flag.data.project.organization);
                $("#project-detail-budget").html(flag.data.project.budget);
                $("#project-detail-createtime").html(flag.data.project.create_time);
                $("#project-detail-inviteway").html(flag.data.project.invite_way);
                $("#project-detail-instruction").html(flag.data.project.instruction);
                $("#teacher-detail-name").html(flag.data.teacher.name);
                $("#teacher-detail-id").html(flag.data.teacher.id);
                if(flag.data.teacher.sex == "f"){
                  $("#teacher-detail-sex").html("女");
                } else {
                  $("#teacher-detail-sex").html("男");
                }
                $("#teacher-detail-organization").html(flag.data.teacher.organization);
                $("#teacher-detail-pt").html(flag.data.teacher.professional_title);
                $("#teacher-detail-tel").html(flag.data.teacher.tel);
                $("#organization-verify-instruction").val(flag.data.project.o_audit_instruction)
                $("#project-detail-purpose").html(flag.data.project.purpose);
                $("#project-detail-function").html(flag.data.project.p_function);
                $("#project-detail-result").html(flag.data.project.expect_result);
              } 
            $("#modal-apply-project-detail").modal('show');
          },

            "click #applyTableAuditing":function(e,value,row,index){
              if (row.status == "专家论证"){
                alert("请进行终审！");
                return;
              }
              last_clicked_apply_project_id = row.id;

              $("#modal-apply-project-Auditing").modal('show');
            },

            "click #applyTableFinAuditing":function(e,value,row,index){
              last_clicked_fin_project_id = row.id;
              if (row.status == "网信中心审核"){
                alert("请进行初审！");
                return;
              }
              var flag;
              $.ajax({
                url: "http://localhost:8080/project/imanager/master/audit",
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
              if (flag.msg == "notReady"){
                alert("请等待所有专家审核完成后再进行终审！");
                return;
              } else if (flag.msg != "success"){
                alert(flag.msg);
                return;
              }

              for (var i=0;i<flag.data.length;i++){
                var result = new String();
                if (flag.data[i].result == "pass"){
                  result = "通过";
                } else {
                  result = "否决";
                }
                var info = "<div class='form-group'>\
                <label>专家:</label>"+"<label>"+flag.data[i].master_id+" "+flag.data[i].master_name+"</label>\
                <br><label>审核意见:</label>"+result+"&nbsp; &nbsp;<label>指导预算:</label>"+"\
                <label>"+flag.data[i].fin_funds+"</label><br>\
                <label>审核说明:</label> <label>"+flag.data[i].instruction+"</label>\
                </div> <hr>";
                $("#fin-modal-body").append(info);
              }
              $("#modal-final-auditing").modal('show');
            }
        }
    
        function AddApplyTableFuncAlty(value,row,index){
          return[
            '<button id="applyTableDetail" type="button" class="btn btn-default">详情</button> &nbsp',
            '<button id="applyTableAuditing" type="button" class="btn btn-default">初审</button> &nbsp',
            '<button id="applyTableFinAuditing" type="button" class="btn btn-default">终审</button>',
          ].join("")
        }
  
          $('#ApplyProjectTable').bootstrapTable({
              url: 'http://localhost:8080/project/imanager/project/apply',         //请求后台的URL（*）
              method: 'get',                      //请求方式（*）
              striped: true,                      //是否显示行间隔色
              cache: false,                       //是否使用缓存，默认为true，所以一般情况下需要设置一下这个属性（*）
              pagination: true,                   //是否显示分页（*）
              sortable: true,                     //是否启用排序
              sortOrder: "asc",                   //排序方式
              sidePagination: "client",           //分页方式：client客户端分页，server服务端分页（*）
              pageNumber: 1,                       //初始化加载第一页，默认第一页
              pageSize: 10,                       //每页的记录行数（*）
              pageList: [10, 15, 20,50],        //可供选择的每页的行数（*）
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
              },{
                field: 'name',
                title: '名称'
              }, {
                field: 'organization',
                title: '所属学院'
              },{
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
  
      $("#btn-choice-master").click(function(){
        $("#dropdown-master-choice").empty();
        var flag;
              $.ajax({
              url: hostip+"project/master/list",
              type: "GET",
              async: false,
              dataType : "JSON",
              success: function(data) {
                      flag = data;
                    },
              error: function (jqXHR) { 
                flag = jqXHR.responseJSON;
              }
              },);
              if (flag.msg == "success"){
                for (var i=0;i<flag.data.length;i++){
                  var info="<a class='dropdown-item dropdown-item-master-choice' onclick='choiceMaster(this)'>"+flag.data[i].id+" "+flag.data[i].name+"</a>";
                  $("#dropdown-master-choice").append(info);
                }
              }
      });

      $("#btn-skip-audit").click(function(){
        if (last_clicked_apply_project_status == "专家论证"){
          alert("此阶段请进行最终审核！");
          return;
        }
        $("#modal-apply-project-Auditing").modal('show');
      });

      //终审通过
      $('#fin-auditing-pass').click(function(){
        var verifyInstruction = $("#fin-verify-instruction").val();
        if (verifyInstruction == "") {
          alert("审核说明不能为空！");
          return;
        }
        var fin_funds = $("#fin-verify-funds").val();
        var r=/^[1-9][0-9]+$/gi;
        if (!r.test(fin_funds)){
          alert("确认的项目资金应该为纯数字！");
          return;
        }

        var flag;
        $.ajax({
            url: "http://localhost:8080/project/imanager/finaudit/pass",
            type: "POST",
            async: false,
            dataType : "JSON",
            data: {"instruction": verifyInstruction,"id":last_clicked_fin_project_id,"funds":fin_funds},
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
        } else {
          $("#ApplyProjectTable").bootstrapTable('refresh');
          $("#modal-final-auditing").modal('hide');
          return;
        }
      });

      //终审不过
      $('#fin-auditing-fail').click(function(){
        var verifyInstruction = $("#fin-verify-instruction").val();
        if (verifyInstruction == "") {
          alert("审核说明不能为空！");
          return;
        }

        var flag;
        $.ajax({
            url: "http://localhost:8080/project/imanager/finaudit/fail",
            type: "POST",
            async: false,
            dataType : "JSON",
            data: {"instruction": verifyInstruction, "id":last_clicked_fin_project_id},
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
        } else {
          $("#ApplyProjectTable").bootstrapTable('refresh');
          $("#modal-final-auditing").modal('hide');
          return;
        }
      });


      //审核申请修改招标项目table
      function Table_ApplyChangeInviteProjectTable_Init() {
        //得到查询的参数
        queryParams = function (params) {
          var temp = {   //这里的键的名字和控制器的变量名必须一直，这边改动，控制器也需要改成一样的
          limit: params.limit,   //页面大小
          offset:params.offset
          };
         return temp;
        }
  
        //Table中按钮绑定事件
        window.applyOperateEvents = {
          "click #applyChangeProjectAudit":function(e,value,row,index){
            last_clicked_change_project_id = row.id;

            $("#modal-change-project-auditing").modal('show');
          },
        }
    
        function AddApplyTableFuncAlty(value,row,index){
          return[
            '<button id="applyChangeProjectAudit" type="button" class="btn btn-default">审核</button> ',
          ].join("")
        }
  
          $('#ApplyChangeInviteProjectTable').bootstrapTable({
              url: 'http://localhost:8080/project/imanager/project/invite/change',         //请求后台的URL（*）
              method: 'get',                      //请求方式（*）
              striped: true,                      //是否显示行间隔色
              cache: false,                       //是否使用缓存，默认为true，所以一般情况下需要设置一下这个属性（*）
              pagination: true,                   //是否显示分页（*）
              sortable: true,                     //是否启用排序
              sortOrder: "asc",                   //排序方式
              sidePagination: "client",           //分页方式：client客户端分页，server服务端分页（*）
              pageNumber: 1,                       //初始化加载第一页，默认第一页
              pageSize: 10,                       //每页的记录行数（*）
              pageList: [5, 10, 25, 50, 100],        //可供选择的每页的行数（*）
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
              },{
                field: 'name',
                title: '名称'
              }, {
                field: 'invite_way',
                title: '招标方式',
                sortable: true
              },{
                field: 'funds',
                title: '计划预算'
              },{
                field: 'change_reason',
                title: '申请修改说明'
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

      //修改申请项目通过
      $('#change-project-pass').click(function(){
        var flag;
        $.ajax({
            url: "http://localhost:8080/project/imanager/project/invite/change/pass",
            type: "POST",
            async: false,
            dataType : "JSON",
            data: {"id":last_clicked_change_project_id},
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
        } else {
          $("#ApplyChangeInviteProjectTable").bootstrapTable('refresh');
          $("#modal-change-project-auditing").modal('hide');
          return;
        }
      });

      //修改申请项目不过
      $('#change-project-fail').click(function(){
        var flag;
        $.ajax({
            url: "http://localhost:8080/project/imanager/project/invite/change/fail",
            type: "POST",
            async: false,
            dataType : "JSON",
            data: {"id":last_clicked_change_project_id},
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
        } else {
          $("#ApplyChangeInviteProjectTable").bootstrapTable('refresh');
          $("#modal-change-project-auditing").modal('hide');
          return;
        }
      });


  });