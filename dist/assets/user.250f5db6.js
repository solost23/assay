import{I as s,J as a}from"./index.1248482e.js";const t=async a=>await s({url:"/system/user/list",method:"get",params:a}),e=async t=>await s({url:"/system/user/"+a(t),method:"get"}),r=async a=>await s({url:"/system/user",method:"post",data:a}),u=async a=>await s({url:"/system/user",method:"put",data:a}),d=async a=>await s({url:"/system/user/"+a,method:"delete"}),m=async(a,t)=>await s({url:"/system/user/resetPwd",method:"put",data:{userId:a,password:t}}),o=async(a,t)=>await s({url:"/system/user/changeStatus",method:"put",data:{userId:a,status:t}}),y=async()=>await s({url:"/system/user/profile",method:"get"}),l=async a=>await s({url:"/system/user/profile",method:"put",data:a}),i=async(a,t)=>await s({url:"/system/user/profile/updatePwd",method:"put",params:{oldPassword:a,newPassword:t}}),p=async a=>await s({url:"/system/user/profile/avatar",method:"post",data:a}),w=async a=>await s({url:"/system/user/authRole/"+a,method:"get"}),h=async a=>await s({url:"/system/user/authRole",method:"put",params:a});export{l as a,i as b,w as c,h as d,u as e,r as f,y as g,d as h,o as i,e as j,t as l,m as r,p as u};