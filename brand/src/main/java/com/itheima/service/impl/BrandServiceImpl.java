package com.itheima.service.impl;

import com.itheima.mapper.BrandMapper;
import com.itheima.pojo.Brand;
import com.itheima.pojo.PageBean;
import com.itheima.service.BrandService;
import com.itheima.util.SqlSessionFactoryUtils;
import org.apache.ibatis.session.SqlSession;
import org.apache.ibatis.session.SqlSessionFactory;

import java.util.List;

public class BrandServiceImpl implements BrandService {
    //1.创建sqlSessionFactory 工厂对象
    SqlSessionFactory factory = SqlSessionFactoryUtils.getSqlSessionFactory();

    public List<Brand> selectAll() {
        //获取SqlSession对象
        SqlSession sqlSession = factory.openSession();
        //获取BrandMapper
        BrandMapper mapper = sqlSession.getMapper(BrandMapper.class);

        //调用方法
        List<Brand> brands = mapper.selectAll();

        //释放资源
        sqlSession.close();

        return brands;
    }

    public void add(Brand brand) {
        //获取SqlSession对象
        SqlSession sqlSession = factory.openSession();
        //获取BrandMapper
        BrandMapper mapper = sqlSession.getMapper(BrandMapper.class);

        //调用方法
        mapper.add(brand);
        sqlSession.commit(); //提交事务

        //释放资源
        sqlSession.close();

    }

    public void update(Brand brand) {
        //获取SqlSession对象
        SqlSession sqlSession = factory.openSession();
        //获取BrandMapper
        BrandMapper mapper = sqlSession.getMapper(BrandMapper.class);

        //调用方法
        mapper.update(brand);
        sqlSession.commit(); //提交事务

        //释放资源
        sqlSession.close();

    }

    public Brand selectById(int id) {
        //获取SqlSession对象
        SqlSession sqlSession = factory.openSession();
        //获取BrandMapper
        BrandMapper mapper = sqlSession.getMapper(BrandMapper.class);

        //调用方法
        Brand brand = mapper.selectById(id);

        //释放资源
        sqlSession.close();

        return brand;
    }

    public void deleteById(int id) {
        //获取SqlSession对象
        SqlSession sqlSession = factory.openSession();
        //获取BrandMapper
        BrandMapper mapper = sqlSession.getMapper(BrandMapper.class);

        //调用方法
        mapper.deleteById(id);
        sqlSession.commit();

        //释放资源
        sqlSession.close();
    }

    public void deleteByIds(int[] id) {
        //获取SqlSession对象
        SqlSession sqlSession = factory.openSession();
        //获取BrandMapper
        BrandMapper mapper = sqlSession.getMapper(BrandMapper.class);

        //调用方法
        mapper.deleteByIds(id);
        sqlSession.commit();

        //释放资源
        sqlSession.close();

    }

    public PageBean<Brand> selectByPageAndCondition(int currentPage, int pageSize, Brand brand) {
        //获取SqlSession对象
        SqlSession sqlSession = factory.openSession();
        //获取BrandMapper
        BrandMapper mapper = sqlSession.getMapper(BrandMapper.class);

        PageBean<Brand> pageBean = new PageBean<Brand>();
        //计算开始索引
        int begin = (currentPage - 1) * pageSize;
        int size = pageSize;
        //处理brand,模糊表达式
        String brandName = brand.getBrandName();
        if(brandName != null && brandName.length() > 0){
            brand.setBrandName("%" + brandName + "%");
        }
        String companyName = brand.getCompanyName();
        if(companyName != null && companyName.length() > 0){
            brand.setCompanyName("%" + companyName + "%");
        }
        //调用方法，封装对象
        pageBean.setRows(mapper.selectByPageAndCondition(begin, size,brand));
        pageBean.setTotalCount(mapper.selectTotalByCondition(brand));

        //释放资源
        sqlSession.close();
        return pageBean;
    }

    public PageBean<Brand> selectByPage(int currentPage, int pageSize) {
        //获取SqlSession对象
        SqlSession sqlSession = factory.openSession();
        //获取BrandMapper
        BrandMapper mapper = sqlSession.getMapper(BrandMapper.class);

        PageBean<Brand> pageBean = new PageBean<Brand>();
        //计算开始索引
        int begin = (currentPage - 1) * pageSize;
        int size = pageSize;
        //调用方法，封装对象
        List<Brand> brands = mapper.selectByPage(begin, size);
        pageBean.setRows(brands);
        pageBean.setTotalCount(mapper.selectTotal());

        //释放资源
        sqlSession.close();
        return pageBean;

    }
}
