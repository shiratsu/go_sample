CREATE TABLE `company_info` (
  `id` int(11) NOT NULL auto_increment,
  `c_name` varchar(64) default NULL,
  `star` varchar(3) default NULL,
  `url` varchar(300) default NULL,
  `money_star` varchar(3) default NULL,
  `grow_star` varchar(3) default NULL,
  `policy_star` varchar(3) default NULL,
  `comfort_star` varchar(3) default NULL,
  `fukurikousei_star` varchar(3) default NULL,
  `safe_star` varchar(3) default NULL,
  `yarigai_star` varchar(3) default NULL,
  `brand_star` varchar(3) default NULL,
  `entry_star` varchar(3) default NULL,
  `education_star` varchar(3) default NULL,
  PRIMARY KEY  (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8