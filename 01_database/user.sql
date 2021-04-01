CREATE TABLE `user` (
  `ID` int(10) UNSIGNED NOT NULL,
  `UserName` varchar(16) NOT NULL,
  `Parent` int(10) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `user` (`ID`, `UserName`, `Parent`) VALUES
(1, 'Ali', 2),
(2, 'Budi', NULL),
(3, 'Cecep', 1);

ALTER TABLE `user`
  ADD PRIMARY KEY (`ID`),
  ADD KEY `Parent` (`Parent`);

ALTER TABLE `user`
  MODIFY `ID` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

ALTER TABLE `user`
  ADD CONSTRAINT `Parent` FOREIGN KEY (`Parent`) REFERENCES `user` (`ID`);
