INSERT INTO `todo` (`id`, `title`) VALUES ('ID1', 'TitleA')
,('ID2', 'TitleB')
,('ID3', 'TitleC')
,('ID4', 'TitleD')
ON DUPLICATE KEY UPDATE
	`id` = VALUES(`id`)
;