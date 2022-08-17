package com.itheima.service;

import com.itheima.pojo.Brand;
import com.itheima.pojo.PageBean;

import java.util.List;

public interface BrandService {
    /**
     * 查询所有
     * @return
     */
    List<Brand> selectAll();

    /**
     * 添加数据
     * @param brand
     */
    void add(Brand brand);

    /**
     * 修改数据
     * @param brand
     */
    void update(Brand brand);
    /**
     * 通过 id 查询数据
     * @param id
     */
    Brand selectById(int id);

    /**
     * 通过id 删除数据
     * @param id
     */
    void deleteById(int id);
    /**
     * 通过id 批量删除数据
     * @param id
     */
    void deleteByIds(int[] id);
    /**
     * 获得分页查询数据 和 总数据条数
     * @param currentPage 当前页码
     * @param pageSize 每页展示条数
     */
    PageBean<Brand> selectByPage(int currentPage, int pageSize);
    /**
     * 获得 分页条件 查询数据 和 总数据条数
     * @param currentPage 当前页码
     * @param pageSize 每页展示条数
     * @param brand 条件
     */
    PageBean<Brand> selectByPageAndCondition(int currentPage, int pageSize, Brand brand);
}
