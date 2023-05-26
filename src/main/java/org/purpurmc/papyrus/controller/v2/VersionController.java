package org.purpurmc.papyrus.controller.v2;

import org.purpurmc.papyrus.db.entity.Build;
import org.purpurmc.papyrus.db.entity.Project;
import org.purpurmc.papyrus.db.entity.Version;
import org.purpurmc.papyrus.db.repository.BuildRepository;
import org.purpurmc.papyrus.db.repository.ProjectRepository;
import org.purpurmc.papyrus.db.repository.VersionRepository;
import org.purpurmc.papyrus.exception.ProjectNotFound;
import org.purpurmc.papyrus.exception.VersionNotFound;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;
import java.util.Optional;

@RestController
@RequestMapping("/v2/{project}")
public class VersionController {
    private final ProjectRepository projectRepository;
    private final VersionRepository versionRepository;
    private final BuildRepository buildRepository;

    @Autowired
    public VersionController(ProjectRepository projectRepository, VersionRepository versionRepository, BuildRepository buildRepository) {
        this.projectRepository = projectRepository;
        this.versionRepository = versionRepository;
        this.buildRepository = buildRepository;
    }

    @GetMapping("/{version}")
    @ResponseBody
    public GetVersion getVersion(@PathVariable("project") String projectName, @PathVariable("version") String versionName) {
        Project project = projectRepository.findByName(projectName).orElseThrow(ProjectNotFound::new);
        Version version = versionRepository.findByProjectAndName(project, versionName).orElseThrow(VersionNotFound::new);
        List<Build> builds = buildRepository.findAllByVersionAndFileNotNullOrderByTimestampAsc(version);
        Optional<Build> latest = buildRepository.findLatestByVersionAndFileNotNull(version);

        return new GetVersion(project.getName(), version.getName(), new GetVersion.Builds(latest.map(Build::getName), builds.stream().map(Build::getName).toList()));
    }

    private record GetVersion(String project, String version, Builds builds) {
        public record Builds(Optional<String> latest, List<String> all) {
        }
    }
}
