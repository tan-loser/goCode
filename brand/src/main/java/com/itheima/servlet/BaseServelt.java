package com.itheima.servlet;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;

/**
 * 替换 HttpServlet 通过请求路径的最后一个路径分发
 */
public class BaseServelt extends HttpServlet {

    @Override
    public void service(HttpServletRequest req, HttpServletResponse res) throws ServletException, IOException {
        //获取请求路径
        String uri = req.getRequestURI();
        //获取最后一段路径，方法名
        int index = uri.lastIndexOf("/");
        String methodName = uri.substring(index+1); //  带斜杠的方法名，需要+1跳过斜杠

        //执行方法 通过反射
        //获取class对象
        //this 代表 brandServlet，不是BaseServlet,谁调用方法，this代表谁
        Class<? extends BaseServelt> cls = this.getClass();
        //获取方法对象
        try {
            Method method = cls.getMethod(methodName, HttpServletRequest.class, HttpServletResponse.class);
            //执行方法
            method.invoke(this,req,res);
        } catch (NoSuchMethodException e) {
            e.printStackTrace();
        } catch (IllegalAccessException e) {
            e.printStackTrace();
        } catch (InvocationTargetException e) {
            e.printStackTrace();
        }


    }
}
