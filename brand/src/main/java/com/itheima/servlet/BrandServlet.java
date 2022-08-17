package com.itheima.servlet;

import com.alibaba.fastjson.JSON;
import com.itheima.pojo.Brand;
import com.itheima.pojo.PageBean;
import com.itheima.service.BrandService;
import com.itheima.service.impl.BrandServiceImpl;

import javax.servlet.ServletException;
import javax.servlet.annotation.WebServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.BufferedReader;
import java.io.IOException;
import java.util.List;

@WebServlet("/brand/*")
public class BrandServlet extends BaseServelt{
    private BrandService brandService = new BrandServiceImpl();

    public void selectAll(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
        //调用service查询
        List<Brand> brands = brandService.selectAll();

        //转换为json
        String jsonString = JSON.toJSONString(brands);

        //写数据
        response.setContentType("text/json;charset=utf-8");
        response.getWriter().write(jsonString);
    }
    public void add(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
        //接受品牌数据
        BufferedReader reader = request.getReader();
        String str = reader.readLine(); //json 字符串

        //转换成Brand对象
        Brand brand = JSON.parseObject(str, Brand.class);

        //调用service添加数据
        brandService.add(brand);

        //响应成功的标志
        response.getWriter().write("success");
    }
    public void update(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
        //接受品牌数据
        BufferedReader reader = request.getReader();
        String str = reader.readLine(); //json 字符串

        //转换成Brand对象
        Brand brand = JSON.parseObject(str, Brand.class);

        //调用service添加数据
        brandService.update(brand);

        //响应成功的标志
        response.getWriter().write("success");
    }
    public void selectById(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
        //接受id
        int id = Integer.parseInt(request.getParameter("id"));

        //调用service添加数据
        Brand brand = brandService.selectById(id);

        //转换为json
        String jsonString = JSON.toJSONString(brand);

        //写数据
        response.setContentType("text/json;charset=utf-8");
        response.getWriter().write(jsonString);
    }
    public void deleteById(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
        //接受id
        int id = Integer.parseInt(request.getParameter("id"));

        //调用service添加数据
        brandService.deleteById(id);

        response.getWriter().write("success");
    }
    public void deleteByIds(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
        //接受id数组json
        BufferedReader reader = request.getReader();
        String str = reader.readLine(); //json 字符串

        //转换成 数组
        int[] ids = JSON.parseObject(str, int[].class);

        //调用service添加批量删除
        brandService.deleteByIds(ids);

        response.getWriter().write("success");
    }
    public void selectByPage(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
        //获取当前页码 和 每页展示条数
        int currentPage = Integer.parseInt(request.getParameter("currentPage"));
        int pageSize = Integer.parseInt(request.getParameter("pageSize"));

        //调用service 查询
        PageBean<Brand> pageBean = brandService.selectByPage(currentPage, pageSize);

        //转换为JSON
        String jsonString = JSON.toJSONString(pageBean);

        //返回数据
        response.setContentType("text/json;charset=utf-8");
        response.getWriter().write(jsonString);
    }
    public void selectByPageAndCondition(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
        //获取当前页码 和 每页展示条数
        int currentPage = Integer.parseInt(request.getParameter("currentPage"));
        int pageSize = Integer.parseInt(request.getParameter("pageSize"));
        //接受查询条件
        BufferedReader reader = request.getReader();
        String str = reader.readLine(); //json 字符串

        //转换成对象
        Brand brand = JSON.parseObject(str, Brand.class);

        //调用service 查询
        PageBean<Brand> pageBean = brandService.selectByPageAndCondition(currentPage, pageSize,brand);

        //转换为JSON
        String jsonString = JSON.toJSONString(pageBean);

        //返回数据
        response.setContentType("text/json;charset=utf-8");
        response.getWriter().write(jsonString);
    }

}
