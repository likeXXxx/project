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
            },
        }

        function AddAbolitionTableFuncAlty(value,row,index){
            return[
              '<button id="finished-project-detail" type="button" class="btn btn-default">项目详情</button>',
              '<button id="finished-project-invite" type="button" class="btn btn-default">招标信息</button>',
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
});