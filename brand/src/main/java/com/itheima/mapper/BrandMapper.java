package com.itheima.mapper;

import com.itheima.pojo.Brand;
import org.apache.ibatis.annotations.*;

import java.util.List;

public interface BrandMapper {

    /**
     * 查询所有
     * @return
     */
    @Select("select * from tb_brand")
    @ResultMap("brandMapper")
    List<Brand> selectAll();

    /**
     * 添加数据
     * @param brand
     */
    @Insert("insert into tb_brand values(null,#{brandName},#{companyName},#{ordered},#{description},#{status})")
    void add(Brand brand);

    /**
     * 修改
     * @param brand
     */
    @Update("update tb_brand set brand_name = #{brandName},company_name = #{companyName},ordered = #{ordered},description = #{description},status = #{status} where id = #{id}")
    void update(Brand brand);
    /**
     * 通过 id 查询数据
     * @param id
     */
    @Select("select * from tb_brand where id = #{id}")
    @ResultMap("brandMapper")
    Brand selectById(int id);

    /**
     * 通过id删除数据
     * @param id
     */
    @Delete("delete from tb_brand where id = #{id}")
    void deleteById(int id);

    /**
     * 通过id 批量删除数据
     * @param id
     */
    void deleteByIds(@Param("ids") int[] id);

    /**
     * 获得分页查询数据
     * @param begin
     * @param size
     */
    @Select("select * from tb_brand limit #{begin}, #{size}")
    @ResultMap("brandMapper")
    List<Brand> selectByPage(@Param("begin") int begin, @Param("size") int size);

    /**
     * 获得分页查询的总数据条数
     * @return
     */
    @Select("select count(*) from tb_brand")
    int selectTotal();

    /**
     * 获得 分页 条件 查询数据
     * @param begin
     * @param size
     * @param brand
     */
    List<Brand> selectByPageAndCondition(@Param("begin") int begin, @Param("size") int size, @Param("brand") Brand brand);

    /**
     * 获得 分页条件 查询的总数据条数
     * @return
     */
    int selectTotalByCondition(Brand brand);

}
