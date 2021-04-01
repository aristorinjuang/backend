CREATE TABLE `logs` (
  `id` int(10) UNSIGNED NOT NULL,
  `searchword` tinytext NOT NULL,
  `pagination` int(10) UNSIGNED NOT NULL,
  `response` text NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE `logs`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `logs`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;