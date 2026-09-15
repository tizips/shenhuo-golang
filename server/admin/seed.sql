-- =====================================================================
-- 神火 - 业务表测试数据（直接操作数据库）
-- 用法：mysql -h <host> -P <port> -u root -p shenhuo < seed.sql
-- 人员测试密码统一为：123456
-- =====================================================================
SET NAMES utf8mb4;

-- 清空业务表（不含账号/角色/配置）
DELETE FROM `sh_score`;
DELETE FROM `sh_draw`;
DELETE FROM `sh_draw_category`;
DELETE FROM `sh_media`;
DELETE FROM `sh_scene`;
DELETE FROM `sh_person`;
DELETE FROM `sh_article`;
DELETE FROM `sh_page`;
DELETE FROM `sh_nav`;
DELETE FROM `sh_schedule`;
DELETE FROM `sh_schedule_category`;
DELETE FROM `sh_banner`;

-- ---------------------------------------------------------------------
-- 人员（密码均为 123456）
-- ---------------------------------------------------------------------
INSERT INTO `sh_person` (`id`, `name`, `unit`, `mobile`, `password`, `number`, `group_name`, `must_change_password`) VALUES
('1970000000000000001', '张伟', '神火集团有限公司', '13800000001', '$2a$10$pjNcf5zZ4DMWsTmK6tLX0eDy8GL0BMmKmRZUBc8hSL4MSzSVYodm2', '001', '第一组', 2),
('1970000000000000002', '李娜', '神火集团有限公司', '13800000002', '$2a$10$pjNcf5zZ4DMWsTmK6tLX0eDy8GL0BMmKmRZUBc8hSL4MSzSVYodm2', '002', '第一组', 2),
('1970000000000000003', '王强', '煤炭实业分公司',   '13800000003', '$2a$10$pjNcf5zZ4DMWsTmK6tLX0eDy8GL0BMmKmRZUBc8hSL4MSzSVYodm2', '003', '第一组', 2),
('1970000000000000004', '刘敏', '煤炭实业分公司',   '13800000004', '$2a$10$pjNcf5zZ4DMWsTmK6tLX0eDy8GL0BMmKmRZUBc8hSL4MSzSVYodm2', '004', '第二组', 2),
('1970000000000000005', '陈杰', '电力运营中心',     '13800000005', '$2a$10$pjNcf5zZ4DMWsTmK6tLX0eDy8GL0BMmKmRZUBc8hSL4MSzSVYodm2', '005', '第二组', 2),
('1970000000000000006', '赵磊', '电力运营中心',     '13800000006', '$2a$10$pjNcf5zZ4DMWsTmK6tLX0eDy8GL0BMmKmRZUBc8hSL4MSzSVYodm2', '006', '第三组', 2),
('1970000000000000007', '孙丽', '化工新材料公司',   '13800000007', '$2a$10$pjNcf5zZ4DMWsTmK6tLX0eDy8GL0BMmKmRZUBc8hSL4MSzSVYodm2', '007', '第三组', 2),
('1970000000000000008', '周涛', '化工新材料公司',   '13800000008', '$2a$10$pjNcf5zZ4DMWsTmK6tLX0eDy8GL0BMmKmRZUBc8hSL4MSzSVYodm2', '008', '第四组', 2);

-- ---------------------------------------------------------------------
-- 抽奖分类
-- ---------------------------------------------------------------------
INSERT INTO `sh_draw_category` (`id`, `name`, `icon`, `quota`, `order`) VALUES
(1, '一等奖', 'https://picsum.photos/seed/shenhuo-draw-1st/750/420', 2,  10),
(2, '二等奖', 'https://picsum.photos/seed/shenhuo-draw-2nd/750/420', 5,  20),
(3, '三等奖', 'https://picsum.photos/seed/shenhuo-draw-3rd/750/420', 10, 30);

-- ---------------------------------------------------------------------
-- 抽奖记录
-- ---------------------------------------------------------------------
INSERT INTO `sh_draw` (`category_id`, `person_id`, `created_at`) VALUES
(1, '1970000000000000001', NOW()),
(1, '1970000000000000005', NOW()),
(2, '1970000000000000002', NOW()),
(2, '1970000000000000004', NOW()),
(2, '1970000000000000007', NOW()),
(3, '1970000000000000003', NOW()),
(3, '1970000000000000006', NOW()),
(3, '1970000000000000008', NOW());

-- ---------------------------------------------------------------------
-- 场景与媒体
-- ---------------------------------------------------------------------
INSERT INTO `sh_scene` (`id`, `name`, `order`) VALUES
(1, '开幕式',   10),
(2, '比赛现场', 20),
(3, '闭幕式',   30);

INSERT INTO `sh_media` (`scene_id`, `type`, `title`, `url`, `is_top`) VALUES
(1, 'image', '开幕式全景', 'https://picsum.photos/seed/shenhuo-scene-1-1/750/420', 1),
(1, 'image', '方队入场',   'https://picsum.photos/seed/shenhuo-scene-1-2/750/420', 2),
(1, 'video', '开幕式视频', 'https://www.w3schools.com/html/mov_bbb.mp4',           2),
(2, 'image', '百米冲刺',   'https://picsum.photos/seed/shenhuo-scene-2-1/750/420', 1),
(2, 'image', '跳远瞬间',   'https://picsum.photos/seed/shenhuo-scene-2-2/750/420', 2),
(2, 'image', '拔河对决',   'https://picsum.photos/seed/shenhuo-scene-2-3/750/420', 2),
(3, 'image', '颁奖合影',   'https://picsum.photos/seed/shenhuo-scene-3-1/750/420', 1),
(3, 'video', '闭幕式视频', 'https://www.w3schools.com/html/mov_bbb.mp4',           2);

-- ---------------------------------------------------------------------
-- 单页（富文本含图片/视频）
-- ---------------------------------------------------------------------
INSERT INTO `sh_page` (`id`, `title`, `content`) VALUES
(1, '关于我们', '<p>神火集团是集煤焦、电力、化工于一体的大型企业集团，拥有员工数万人，多年位居中国企业500强前列。</p><p style="text-align:center;"><img src="https://picsum.photos/seed/shenhuo-page-about-1/750/420" alt="集团园区" style="max-width:100%;border-radius:8px;"/></p><p>集团始终坚持“安全第一、绿色发展”的理念，持续推进产业升级与技术创新。</p><p style="text-align:center;"><img src="https://picsum.photos/seed/shenhuo-page-about-2/750/420" alt="厂区全景" style="max-width:100%;border-radius:8px;"/></p>'),
(2, '参赛须知', '<p>请各位运动员提前30分钟到场检录，注意人身安全。</p><p style="text-align:center;"><img src="https://picsum.photos/seed/shenhuo-page-notice-1/750/420" alt="检录处示意图" style="max-width:100%;border-radius:8px;"/></p><p>检录流程及注意事项请观看下方视频：</p><p style="text-align:center;"><video src="https://www.w3schools.com/html/mov_bbb.mp4" controls poster="https://picsum.photos/seed/shenhuo-page-notice-video/750/420" style="max-width:100%;border-radius:8px;"></video></p><p>比赛期间请穿着运动服与运动鞋，服从现场裁判与工作人员安排，如有身体不适请立即停止比赛并联系医护人员。</p>'),
(3, '联系方式', '<p>组委会电话：0375-88888888</p><p>组委会邮箱：games@shenhuo.com</p><p style="text-align:center;"><img src="https://picsum.photos/seed/shenhuo-page-contact-1/750/420" alt="组委会办公地点" style="max-width:100%;border-radius:8px;"/></p>');

-- ---------------------------------------------------------------------
-- 导航（page 类型引用单页 ID）
-- ---------------------------------------------------------------------
INSERT INTO `sh_nav` (`title`, `icon`, `type`, `value`, `order`) VALUES
('赛事日程', 'https://picsum.photos/seed/shenhuo-nav-schedule/750/420', 'link', '/pages/schedule',        10),
('成绩查询', 'https://picsum.photos/seed/shenhuo-nav-score/750/420',     'link', '/pages/score',           20),
('参赛须知', 'https://picsum.photos/seed/shenhuo-nav-notice/750/420',    'page', '2',                      30),
('关于我们', 'https://picsum.photos/seed/shenhuo-nav-about/750/420',     'page', '1',                      40),
('官方网站', 'https://picsum.photos/seed/shenhuo-nav-site/750/420',      'link', 'https://www.shenhuo.com', 50);

-- ---------------------------------------------------------------------
-- 文章（富文本含图片/视频）
-- ---------------------------------------------------------------------
INSERT INTO `sh_article` (`title`, `thumb`, `content`, `published_at`, `is_top`, `is_recommend`) VALUES
('第X届职工运动会隆重开幕', 'https://picsum.photos/seed/shenhuo-article-1/750/420', '<p>9月1日上午，集团第X届职工运动会在体育场隆重开幕，来自各单位的代表队参加了开幕式。</p><p style="text-align:center;"><img src="https://picsum.photos/seed/shenhuo-article-1-a/750/420" alt="开幕式现场" style="max-width:100%;border-radius:8px;"/></p><p>开幕式上，方队整齐入场，展现了职工们昂扬向上的精神风貌。</p><p style="text-align:center;"><img src="https://picsum.photos/seed/shenhuo-article-1-b/750/420" alt="方队入场" style="max-width:100%;border-radius:8px;"/></p><p>精彩瞬间请观看视频：</p><p style="text-align:center;"><video src="https://www.w3schools.com/html/mov_bbb.mp4" controls poster="https://picsum.photos/seed/shenhuo-article-1-video/750/420" style="max-width:100%;border-radius:8px;"></video></p>', '2026-09-01 09:00:00', 1, 1),
('男子100米决赛精彩回顾',   'https://picsum.photos/seed/shenhuo-article-2/750/420', '<p>男子100米决赛中，运动员们奋力拼搏，现场气氛热烈。</p><p style="text-align:center;"><img src="https://picsum.photos/seed/shenhuo-article-2-a/750/420" alt="起跑瞬间" style="max-width:100%;border-radius:8px;"/></p><p>冲线时刻视频回顾：</p><p style="text-align:center;"><video src="https://www.w3schools.com/html/mov_bbb.mp4" controls poster="https://picsum.photos/seed/shenhuo-article-2-video/750/420" style="max-width:100%;border-radius:8px;"></video></p><p>最终成绩将由组委会统一公布，敬请关注。</p>', '2026-09-02 10:30:00', 2, 1),
('拔河比赛圆满收官',         'https://picsum.photos/seed/shenhuo-article-3/750/420', '<p>经过多轮角逐，拔河比赛顺利收官，展现了团队协作精神。</p><p style="text-align:center;"><img src="https://picsum.photos/seed/shenhuo-article-3-a/750/420" alt="拔河对决" style="max-width:100%;border-radius:8px;"/></p>', '2026-09-03 15:00:00', 2, 2),
('运动会期间交通管制提示',   'https://picsum.photos/seed/shenhuo-article-4/750/420', '<p>运动会期间体育场周边道路将实施临时交通管制，请提前规划出行路线。</p><p style="text-align:center;"><img src="https://picsum.photos/seed/shenhuo-article-4-a/750/420" alt="交通管制示意图" style="max-width:100%;border-radius:8px;"/></p>', '2026-09-04 08:00:00', 2, 2),
('运动会摄影作品征集启事',   'https://picsum.photos/seed/shenhuo-article-5/750/420', '<p>现面向全体职工征集运动会摄影作品，优秀作品将在集团展厅展出。</p><p style="text-align:center;"><img src="https://picsum.photos/seed/shenhuo-article-5-a/750/420" alt="往届作品" style="max-width:100%;border-radius:8px;"/></p><p>参考视频：</p><p style="text-align:center;"><video src="https://www.w3schools.com/html/mov_bbb.mp4" controls poster="https://picsum.photos/seed/shenhuo-article-5-video/750/420" style="max-width:100%;border-radius:8px;"></video></p>', '2026-09-05 14:00:00', 2, 1);

-- ---------------------------------------------------------------------
-- 日程分类
-- ---------------------------------------------------------------------
INSERT INTO `sh_schedule_category` (`id`, `name`, `order`) VALUES
(1, '开幕式安排',   10),
(2, '决赛日程安排', 20),
(3, '闭幕式安排',   30);

-- ---------------------------------------------------------------------
-- 赛程
-- ---------------------------------------------------------------------
INSERT INTO `sh_schedule` (`category_id`, `title`, `subtitle`, `description`, `time`, `items`, `order`) VALUES
(1, '开幕式',           '入场仪式、领导致辞、运动员宣誓', '请各单位代表队8:30前到场集合', '09:00', '[{"name":"方队入场","time":"09:00"},{"name":"领导致辞","time":"09:20"},{"name":"运动员宣誓","time":"09:40"}]', 10),
(2, '男子100米预赛',     '分组预赛',                       '每组前两名晋级决赛',           '10:00', '[]', 20),
(2, '女子跳远决赛',     '决赛',                           '每人试跳三次，取最好成绩',     '14:00', '[]', 30),
(2, '拔河比赛',         '小组赛及淘汰赛',                 '各单位选派10名队员参赛',       '15:30', '[{"name":"小组赛","time":"15:30"},{"name":"决赛","time":"16:30"}]', 40),
(3, '闭幕式暨颁奖仪式', '颁奖典礼、文艺演出',             '请获奖代表队提前候场',         '17:00', '[]', 50);

-- ---------------------------------------------------------------------
-- 轮播图
-- ---------------------------------------------------------------------
INSERT INTO `sh_banner` (`title`, `image`, `link`, `order`) VALUES
('运动会开幕啦',   'https://picsum.photos/seed/shenhuo-banner-1/750/420', '',              10),
('赛事日程一览',   'https://picsum.photos/seed/shenhuo-banner-2/750/420', '/pages/schedule', 20),
('成绩实时查询',   'https://picsum.photos/seed/shenhuo-banner-3/750/420', '/pages/score',    30),
('精彩瞬间回顾',   'https://picsum.photos/seed/shenhuo-banner-4/750/420', '/pages/moment',   40);

-- ---------------------------------------------------------------------
-- 成绩（一人一条）
-- ---------------------------------------------------------------------
INSERT INTO `sh_score` (`person_id`, `total`, `items`, `order`) VALUES
('1970000000000000001', '60',  '[{"name":"100米","value":"14.50"},{"name":"跳远","value":"3.80"},{"name":"铅球","value":"6.50"}]', 10),
('1970000000000000002', '65',  '[{"name":"100米","value":"14.70"},{"name":"跳远","value":"3.90"},{"name":"铅球","value":"6.80"}]', 20),
('1970000000000000003', '70',  '[{"name":"100米","value":"14.90"},{"name":"跳远","value":"4.00"},{"name":"铅球","value":"7.10"}]', 30),
('1970000000000000004', '75',  '[{"name":"100米","value":"15.10"},{"name":"跳远","value":"4.10"},{"name":"铅球","value":"7.40"}]', 40),
('1970000000000000005', '80',  '[{"name":"100米","value":"15.30"},{"name":"跳远","value":"4.20"},{"name":"铅球","value":"7.70"}]', 50),
('1970000000000000006', '85',  '[{"name":"100米","value":"15.50"},{"name":"跳远","value":"4.30"},{"name":"铅球","value":"8.00"}]', 60),
('1970000000000000007', '90',  '[{"name":"100米","value":"15.70"},{"name":"跳远","value":"4.40"},{"name":"铅球","value":"8.30"}]', 70),
('1970000000000000008', '95',  '[{"name":"100米","value":"15.90"},{"name":"跳远","value":"4.50"},{"name":"铅球","value":"8.60"}]', 80);
