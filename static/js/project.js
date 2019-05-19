$(document).ready(function(){
    PassProjectTable_Init();
    RunningProjectTable_Init();
    FinishedProjectTable_Init();
    var hostip = "http://localhost:8080/";

    $("#li1").click(function(){
        $("#u-name").html("过审项目")
    });
      
    $("#li2").click(function(){
        $("#u-name").html("执行中项目")
    });
      
    $("#li3").click(function(){
        $("#u-name").html("已完成项目")
    });

    $("#btn-login").click(function(){
        window.open(hostip+"project/login");
    });

    //过审项目
    $("#btn-refresh-PassProjectTable").click(function(){
        $("#PassProjectTable").bootstrapTable('refresh');
    });

    function PassProjectTable_Init(){
        queryParams = function (params) {
            var temp = {   
            limit: params.limit,   //页面大小
            offset:params.offset
            };
           return temp;
        };

        //Table中按钮绑定事件
        window.abolitionOperateEvents = {
            "click #pass-project-detail":function(e,value,row,index){
            },
        }

        function AddAbolitionTableFuncAlty(value,row,index){
            return[
              '<button id="pass-project-detail" type="button" class="btn btn-default">项目详情</button>',
              '<button id="pass-project-invite" type="button" class="btn btn-default">招标信息</button>',
            ].join("")
        };

        $('#PassProjectTable').bootstrapTable({
            url: 'http://localhost:8080/project/global/passproject',         //请求后台的URL（*）
            method: 'get',                      //请求方式（*）
            striped: true,                      //是否显示行间隔色
            cache: false,                       //是否使用缓存，默认为true，所以一般情况下需要设置一下这个属性（*）
            pagination: true,                   //是否显示分页（*）
            sortable: true,                     //是否启用排序
            sortOrder: "asc",                   //排序方式
            sidePagination: 'client',           //分页方式：client客户端分页，server服务端分页（*）
            pageNumber: 1,                       //初始化加载第一页，默认第一页
            pageSize: 20,                       //每页的记录行数（*）
            pageList: [10, 20, 50],        //可供选择的每页的行数（*）
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
            },{
                field: 'name',
                title: '名称'
            },{
                field: 'create_time',
                title: '创建时间',
                sortable: true
            },{
                field: 'organization',
                title: '所属学院',
            },{
                field: 'teacher',
                title: '负责教师',
            },{
                field: 'teacher_tel',
                title: '联系方式',
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
    }

    //执行中项目
    $("#btn-refresh-RunningProjectTable").click(function(){
        $("#RunningProjectTable").bootstrapTable('refresh');
    });

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
            "click #running-project-detail":function(e,value,row,index){
            },
        }

        function AddAbolitionTableFuncAlty(value,row,index){
            return[
              '<button id="running-project-detail" type="button" class="btn btn-default">项目详情</button>',
              '<button id="running-project-invite" type="button" class="btn btn-default">招标信息</button>',
            ].join("")
        };

        $('#RunningProjectTable').bootstrapTable({
            url: 'http://localhost:8080/project/global/runningproject',         //请求后台的URL（*）
            method: 'get',                      //请求方式（*）
            striped: true,                      //是否显示行间隔色
            cache: false,                       //是否使用缓存，默认为true，所以一般情况下需要设置一下这个属性（*）
            pagination: true,                   //是否显示分页（*）
            sortable: true,                     //是否启用排序
            sortOrder: "asc",                   //排序方式
            sidePagination: 'client',           //分页方式：client客户端分页，server服务端分页（*）
            pageNumber: 1,                       //初始化加载第一页，默认第一页
            pageSize: 20,                       //每页的记录行数（*）
            pageList: [10, 20, 50],        //可供选择的每页的行数（*）
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
            },{
                field: 'name',
                title: '名称'
            },{
                field: 'run_time',
                title: '执行时间',
                sortable: true
            },{
                field: 'organization',
                title: '所属学院',
            },{
                field: 'teacher',
                title: '负责教师',
            },{
                field: 'teacher_tel',
                title: '联系方式',
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
    }

    //已完成项目
    $("#btn-refresh-FinishedProjectTable").click(function(){
        $("#FinishedProjectTable").bootstrapTable('refresh');
    });

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
                var flag; 
                $.ajax({ 
                    url : hostip+"project/global/finished/detail", 
                    type: "get",
                    async: false,
                    dataType : "JSON",
                    data: {"id":row.id},
                    success: function(data) {
                        flag = data;
                    },
                    error: function (jqXHR) { 
                        flag = jqXHR.responseJSON;
                    }
                },);
                if (flag.msg == "success"){
                    $("#finished-project-detail-id").html(flag.data.project.id);
                    $("#finished-project-detail-name").html(flag.data.project.name);
                    $("#finished-project-detail-organization").html(flag.data.project.organization);
                    $("#finished-project-detail-finfunds").html(flag.data.project.fin_funds);
                    $("#finished-project-detail-usedfunds").html(flag.data.project.used_funds);
                    $("#finished-project-detail-createtime").html(flag.data.project.create_time);
                    $("#finished-project-detail-runtime").html(flag.data.project.run_time);
                    $("#finished-project-detail-fintime").html(flag.data.project.fin_time);
                    $("#finished-project-detail-instruction").html(flag.data.project.instruction)
                    $("#finished-project-detail-purpose").html(flag.data.project.purpose);
                    $("#finished-project-detail-function").html(flag.data.project.p_function);
                    $("#finished-project-detail-result").html(flag.data.project.expect_result);
                    $("#finished-project-detail-completionstatus").html(flag.data.project.completion_status);
                    $("#finished-project-detail-selfevaluation").html(flag.data.project.self_evaluation);
                    $("#finished-teacher-detail-name").html(flag.data.teacher.name);
                    $("#finished-teacher-detail-id").html(flag.data.teacher.id);
                    $("#finished-teacher-detail-organization").html(flag.data.teacher.organization);
                    $("#finished-teacher-detail-pt").html(flag.data.teacher.professional_title);
                    $("#finished-teacher-detail-tel").html(flag.data.teacher.tel);
                    $("#modal-finished-project-detail").modal("show");
                    return;
                } else {
                    alert(flag.msg);
                    return;
                }
            },

            "click #finished-project-event":function(e,value,row,index){
                var flag;
                $.ajax({
                    url: "http://localhost:8080/project/global/project/eventlist",
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
                    $("#finished-project-event-list").append(info);
                }

                $("#modal-finished-project-listevent").modal("show");
            },
        }

        function AddAbolitionTableFuncAlty(value,row,index){
            return[
              '<button id="finished-project-detail" type="button" class="btn btn-default">项目详情</button>',
              '<button id="finished-project-event" type="button" class="btn btn-default">项目事件</button>',
            ].join("")
        };

        $('#FinishedProjectTable').bootstrapTable({
            url: 'http://localhost:8080/project/global/finishedproject',         //请求后台的URL（*）
            method: 'get',                      //请求方式（*）
            striped: true,                      //是否显示行间隔色
            cache: false,                       //是否使用缓存，默认为true，所以一般情况下需要设置一下这个属性（*）
            pagination: true,                   //是否显示分页（*）
            sortable: true,                     //是否启用排序
            sortOrder: "asc",                   //排序方式
            sidePagination: 'client',           //分页方式：client客户端分页，server服务端分页（*）
            pageNumber: 1,                       //初始化加载第一页，默认第一页
            pageSize: 20,                       //每页的记录行数（*）
            pageList: [10, 20, 50],        //可供选择的每页的行数（*）
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
            },{
                field: 'name',
                title: '名称'
            },{
                field: 'finish_time',
                title: '结束时间',
                sortable: true
            },{
                field: 'organization',
                title: '所属学院',
            },{
                field: 'teacher',
                title: '负责教师',
            },{
                field: 'teacher_tel',
                title: '联系方式',
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
    }

    $('#modal-finished-project-detail').on('hide.bs.modal',function() {
        $("#finished-project-detail-id").html("");
        $("#finished-project-detail-name").html("");
        $("#finished-project-detail-organization").html("");
        $("#finished-project-detail-finfunds").html("");
        $("#finished-project-detail-usedfunds").html("");
        $("#finished-project-detail-createtime").html("");
        $("#finished-project-detail-runtime").html("");
        $("#finished-project-detail-fintime").html("");
        $("#finished-project-detail-instruction").html("")
        $("#finished-project-detail-purpose").html("");
        $("#finished-project-detail-function").html("");
        $("#finished-project-detail-result").html("");
        $("#finished-project-detail-completionstatus").html("");
        $("#finished-project-detail-selfevaluation").html("");
        $("#finished-teacher-detail-name").html("");
        $("#finished-teacher-detail-id").html("");
        $("#finished-teacher-detail-organization").html("");
        $("#finished-teacher-detail-pt").html("");
        $("#finished-teacher-detail-tel").html("");
    })

      //查看事件模态框关闭
    $('#modal-finished-project-listevent').on('hide.bs.modal',function() {
        $("#finished-project-event-list").empty();
    })
});