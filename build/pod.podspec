Pod::Spec.new do |spec|
  spec.name         = 'Geuth'
  spec.version      = '{{.Version}}'
  spec.license      = { :type => 'GNU Lesser General Public License, Version 3.0' }
  spec.homepage     = 'https://github.com/eutherum/eutherum'
  spec.authors      = { {{range .Contributors}}
		'{{.Name}}' => '{{.Email}}',{{end}}
	}
  spec.summary      = 'iOS Eutherum Client'
  spec.source       = { :git => 'https://github.com/eutherum/eutherum.git', :commit => '{{.Commit}}' }

	spec.platform = :ios
  spec.ios.deployment_target  = '9.0'
	spec.ios.vendored_frameworks = 'Frameworks/Geuth.framework'

	spec.prepare_command = <<-CMD
    curl https://geuthstore.blob.core.windows.net/builds/{{.Archive}}.tar.gz | tar -xvz
    mkdir Frameworks
    mv {{.Archive}}/Geuth.framework Frameworks
    rm -rf {{.Archive}}
  CMD
end
